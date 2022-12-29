package linklist

import (
	"container/list"
	"fmt"
	"sync"
)

type ListNode struct {
	data interface{}
	Next *ListNode
}

// func main() {
// 	var head = new(ListNode)
// 	CreateNode(head, 10)
// 	PrintNode("前：", head)
// 	yyy := reverseList(head)
// 	PrintNode("后：", yyy)
// }

// 反转单链表
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode

	for cur != nil {

		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}

func CreateNode(node *ListNode, max int) {
	cur := node
	for i := 1; i < max; i++ {
		cur.Next = &ListNode{}
		cur.Next.data = i
		cur = cur.Next
	}
}

// 打印链表的方法
func PrintNode(info string, node *ListNode) {
	fmt.Print(info)
	for cur := node; cur != nil; cur = cur.Next {
		fmt.Print(cur.data, " ")
	}
	fmt.Println()
}

// LrU
type Lrucache struct {
	size     int
	values   *list.List
	cacheMap map[interface{}]*list.Element
	lock     sync.Mutex
}

func NewLru(size int) *Lrucache {
	val := list.New()
	return &Lrucache{
		size:     size,
		values:   val,
		cacheMap: make(map[interface{}]*list.Element, size),
	}
}

func (l *Lrucache) Put(k, v interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.values.Len() == l.size {
		back := l.values.Back()
		l.values.Remove(back)
		delete(l.cacheMap, back)
	}

	front := l.values.PushFront(v)
	l.cacheMap[k] = front
}

func (l *Lrucache) Get(k interface{}) (interface{}, bool) {
	v, ok := l.cacheMap[k]
	if ok {
		l.values.MoveToFront(v)
		return v, true
	} else {
		return nil, false
	}
}

func (l *Lrucache) Size() int {
	return l.values.Len()
}
func (l *Lrucache) String() {
	for i := l.values.Front(); i != nil; i = i.Next() {
		fmt.Print(i.Value, "\t")
	}
}
func (l *Lrucache) List() []interface{} {
	var data []interface{}
	for i := l.values.Front(); i != nil; i = i.Next() {
		data = append(data, i.Value)
	}
	return data
}

func (l *Lrucache) Clear() {
	l.lock.Lock()
	defer l.lock.Unlock()
	l.values = list.New()
	l.cacheMap = make(map[interface{}]*list.Element, l.size)

}
