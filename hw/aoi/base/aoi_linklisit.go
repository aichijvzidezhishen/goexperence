package base

import "sync"

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
	// listener AOIEvent
	object interface{}
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

	}
	// if  {

	// }
}

// func NewAoiNode(layer, value float64, entity interface{}, left, right *AoiNode) *AoiNode {
// 	return &AoiNode{
// 		Value:   value,
// 		Entity:  entity,
// 		Left:    left,
// 		Right:   right,
// 		Top:     nil,
// 		Down:    nil,
// 		_sync:   sync.WaitGroup{},
// 	}
// }

// func (n *AoiNode) Add(layer, value float64, entity interface{}, left, right *AoiNode) *AoiNode {
// 	n.Value = value
// 	n.Entity = entity
// 	n.Left = left
// 	n.Right = right
// 	n.Down = nil

// 	for ; layer > 1 && rand.Intn(2) == 0; layer-- {
// 		n = n.Down = NewAoiNode(layer-1, value, entity, left, right)
// 	}

// 	return n
// }

// func (n *AoiNode) Remove() {
// 	var tmp *AoiNode
// 	for n.Top != nil {
// 		tmp = n
// 		n = n.Top
// 		n.Down = tmp
// 	}
// }

// func (n *AoiNode) Move(target float64) {
// 	var cur *AoiNode = n

// 	for {
// 		if target > cur.Value {
// 			for cur.Right != nil && target > cur.Right.Value {
// 				cur = cur.Right
// 			}
// 			if cur.Right != nil {
// 				var findNode *AoiNode = cur
// 				for findNode.Right != nil && target > findNode.Right.Value {
// 					findNode = findNode.Right
// 				}
// 				cur.Left = findNode
// 				cur.Right = findNode.Right
// 				if findNode.Right != nil {
// 					findNode.Right.Left = cur
// 				}
// 				findNode.Right = cur
// 			}
// 		} else {
// 			for cur.Left != nil && target < cur.Left.Value {
// 				cur = cur.Left
// 			}
// 			if cur.Left != nil {
// 				var findNode *AoiNode = cur
// 				for findNode.Left != nil && target < findNode.Left.Value {
// 					findNode = findNode.Left
// 				}
// 				cur.Right = findNode
// 				cur.Left = findNode.Left
// 				if findNode.Left != nil {
// 					findNode.Left.Right = cur
// 				}
// 				findNode.Left = cur
// 			}
// 		}

// 		cur.Value = target
// 		if cur.Top == nil {
// 			break
// 		}
// 		cur = cur.Top
// 	}
// }
