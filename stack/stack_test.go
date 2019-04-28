package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	stack := &Stack{}
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	if stack.Top() != 40 {
		t.Error("Stack head should be 40")
	}
}

func TestPop(t *testing.T) {
	stack := &Stack{}
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	stack.Pop()
	val, _ := stack.Pop()
	if val != 30 {
		t.Error("Stack head should be 30")
	}
	stack2 := &Stack{}
	stack2.Push(1)
	stack2.Pop()
	if !stack2.IsEmpty() {
		t.Error("Stack should be empty")
	}
	_, err := stack2.Pop()
	if err == nil {
		t.Error("Pop on empty stack should throw an error")
	}
}
