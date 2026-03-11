package stackalo

import "testing"

func TestMinStack(t *testing.T) {
	// data := []int{2, 0, 3, 0, 2}
	minStack := Constructor()
	minStack.Push(2)
	minStack.Push(0)
	minStack.Push(3)
	minStack.Push(0)
	minStack.Push(2)
}
