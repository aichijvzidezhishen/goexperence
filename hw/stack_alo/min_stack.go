package stackalo

import "math"

type MinStack struct {
	minStack []int
	stack    []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{math.MaxInt64},
	}
}

func (m *MinStack) Push(val int) {
	m.stack = append(m.stack, val)
	m.minStack = append(m.minStack, min(val, m.minStack[len(m.minStack)-1]))
}

func (m *MinStack) Pop() {
	m.stack = m.stack[:len(m.stack)-1]
	m.minStack = m.minStack[:len(m.minStack)-1]
}

func (m *MinStack) Top() int {
	return m.stack[len(m.stack)-1]
}

func (m *MinStack) GetMin() int {
	if len(m.minStack) == 0 {
		return 0

	}
	return m.minStack[len(m.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
