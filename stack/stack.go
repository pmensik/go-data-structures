package stack

import (
	"errors"
)

type node struct {
	prev  *node
	value int
}

// Stack implementatio based on linked list
type Stack struct {
	head *node
}

// IsEmpty checks whether stack contains any elements
func (stack *Stack) IsEmpty() bool {
	return stack.head == nil
}

// Push element on top of the stack
func (stack *Stack) Push(element int) bool {
	if stack.IsEmpty() {
		stack.head = &node{prev: nil, value: element}
	} else {
		n := &node{prev: stack.head, value: element}
		stack.head = n
	}
	return true
}

// Top returns last inserted element into the stack
func (stack *Stack) Top() int {
	return stack.head.value
}

// Pop removes and returns last inserted element from the stack
func (stack *Stack) Pop() (int, error) {
	if stack.IsEmpty() {
		return -1, errors.New("Can't pop from empty stack")
	}
	val := stack.head.value
	stack.head = stack.head.prev
	return val, nil
}
