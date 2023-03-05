package list

import "fmt"

type SingleList struct {
	val  int
	next *SingleList
}

func (l *SingleList) InsertNode(val, pos int) {
	cur := l
	for cur != nil {
		// if cur.val == pos {
		tem := cur.next
		cur.next = &SingleList{
			val:  0,
			next: nil,
		}
		cur.next.val = val
		cur.next.next = tem

		// }
		cur = cur.next
	}
}

func (l *SingleList) DelNode(val int) (res *SingleList) {
	if l.val == val {
		return l.next
	}

	cur := l
	for cur.next != nil {
		if cur.next.val == val {
			if cur.next.next == nil {
				cur.next = nil
				return l
			} else {
				cur.next = cur.next.next
				return l
			}
		}
		cur = cur.next
	}
	return nil
}

func DelStable(pos int) {

	var single SingleList
	single.val = 0
	s := &single
	for i := 1; i < 10; i++ {
		s.InsertNode(i, i)
	}

	single.DelNode(pos)
	cur := &single
	for cur.next != nil {
		fmt.Println(cur.val)
		cur = cur.next
	}
	fmt.Println("---")

}
