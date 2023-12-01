package base

import (
	"fmt"
	"math"
	"sync"
)

const (
	limit    = 0
	maxLayer = 8
)

type AoiNode struct {
	// Value  float64
	// Entity interface{}
	// Left   *AoiNode
	// Right  *AoiNode
	// Top    *AoiNode
	// Down   *AoiNode
	// _sync  sync.WaitGroup
	xPre, xNext *AoiNode
	yPre, yNext *AoiNode
	x, y        float32
	radis       float32
	listener    AoiEvent
	object      interface{}
}

type AoiList struct {
	xList *AoiNode
	yList *AoiNode
	lock  sync.RWMutex
}

func (l *AoiList) Add(node *AoiNode) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.xList == nil || l.yList == nil {
		l.xList = node
		l.yList = node
		return
	}

	var tail *AoiNode
	found := false
	for cur := l.xList; cur != nil; cur = cur.xNext {
		if cur.x > node.x {
			node.xNext = cur
			if cur.xPre != nil {
				node.xPre = cur.xPre
				cur.xPre.xNext = node
			} else {
				l.xList = node
			}
			cur.xPre = node
			found = true
			break
		}
		tail = cur
	}
	if !found && tail != nil {
		tail.xNext = node
		node.xPre = tail
	}

	tail = nil
	found = false
	for cur := l.yList; cur != nil; cur = cur.yNext {
		if cur.y > node.y {
			node.yNext = cur
			if cur.yPre != nil {
				node.yPre = cur.yPre
				cur.yPre.yNext = node
			} else {
				l.yList = node
			}
			cur.yPre = node
			found = true
			break
		}
		tail = cur
	}
	if tail != nil && !found {
		tail.yNext = node
		node.yPre = tail
	}

}

func (l *AoiList) Remove(node *AoiNode) {
	if node == l.xList {
		if node.xNext != nil {
			l.xList.xNext = node.xNext
			if l.xList.xPre != nil {
				l.xList.xPre = nil
			}
		} else {
			l.xList = nil
		}
	} else if node.xPre != nil && node.xNext != nil {
		node.xPre.xNext = node.xNext
		node.xNext.xPre = node.xPre
	} else if node.xPre != nil {
		node.xPre.xNext = nil
	}

	if node == l.yList {
		if node.yNext != nil {
			l.yList.yNext = node.yNext
			if l.yList.yPre != nil {
				l.yList.yPre = nil
			}
		} else {
			l.xList = nil
		}
	} else if node.yPre != nil && node.yNext != nil {
		node.yPre.yNext = node.yNext
		node.yNext.yPre = node.yPre
	} else if node.yPre != nil {
		node.yPre.yNext = nil
	}

	node.xPre = nil
	node.xNext = nil
	node.yPre = nil
	node.yNext = nil
}

func (l *AoiList) findInNeighbors(node *AoiNode) map[*AoiNode]struct{} {
	var neighbors map[*AoiNode]struct{}
	for curnode := node.xNext; curnode != nil; curnode = curnode.xNext {
		if curnode.x-node.x > node.radis {
			break
		}
		if math.Abs(float64(curnode.y-node.y)) <= float64(node.radis) {
			neighbors[node] = struct{}{}
		}

	}
	for curnode := node.xPre; curnode != nil; curnode = curnode.xPre {
		if node.x-curnode.x > node.radis {
			break
		}
		if math.Abs(float64(curnode.y-node.y)) <= float64(node.radis) {
			neighbors[node] = struct{}{}
		}
	}
	return neighbors
}

func (l *AoiList) EnterAoi(node *AoiNode) {
	l.Add(node)
	var nodelist []*AoiNode
	neighbors := l.findInNeighbors(node)
	for neighbor := range neighbors {
		neighbor.listener.OnEnterAoi([]*AoiNode{node})
		nodelist = append(nodelist, node)
	}
	node.listener.OnEnterAoi(nodelist)
}

func (l *AoiList) LeaveAoi(node *AoiNode) {
	l.Remove(node)

}

type AoiEventTrriger struct {
	node *AoiNode
}

func (a *AoiEventTrriger) OnEnterEvent(nodeList []*AoiNode) {
	for _, node := range nodeList {
		fmt.Println("node enter", node.x, node.y)
	}
}
