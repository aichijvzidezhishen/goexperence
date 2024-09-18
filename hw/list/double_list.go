package list

type ListNode struct {
	data interface{}
	pre  *ListNode
	next *ListNode
}

type DoubleList struct {
	head *ListNode
	tail *ListNode
}

func NewDoubleList() *DoubleList {
	return &DoubleList{
		head: nil,
		tail: nil,
	}
}
func (l *DoubleList) Append(data interface{}) {
	node := &ListNode{
		data: data,
		pre:  nil,
		next: nil,
	}
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		node.pre = l.tail
		l.tail = node
	}
}
