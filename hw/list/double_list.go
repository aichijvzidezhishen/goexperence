package list

type Element struct {
	next, prev *Element

	list  *List
	value any
}

type List struct {
	root Element
	len  int
}

// 初始化List结构体
func (l *List) Init() *List {
	// 将root节点的next指针指向自身
	l.root.next = &l.root
	// 将root节点的prev指针指向自身
	l.root.prev = &l.root
	// 将List的长度置为0
	l.len = 0
	return l
}

func New() *List {
	return new(List).Init()
}

func (l *List) Len() int {
	return l.len
}

func (l *List) Front() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

func (l *List) Back() *Element {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

func (l *List) insert(e, at *Element) *Element {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

/* type ListNode struct {
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

} */
