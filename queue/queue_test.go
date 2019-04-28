package queue

import (
	"testing"
)

func TestEnqueue(t *testing.T) {
	stack := &Queue{}
	stack.Enqueue(20)
	stack.Enqueue(30)
	stack.Enqueue(40)
	if stack.Peek() != 40 {
		t.Error("Queue head should be 40")
	}
}

func TestIsEmpty(t *testing.T) {
	stack := &Queue{}
	if !stack.IsEmpty() {
		t.Error("Queue should be empty when created")
	}
	stack.Enqueue(20)
	if stack.IsEmpty() {
		t.Error("Queue should not be empty after enqueueing of an element")
	}
}

func TestDequeue(t *testing.T) {
	stack := &Queue{}
	stack.Enqueue(20)
	stack.Enqueue(30)
	stack.Dequeue()
	if stack.Peek() != 20 {
		t.Error("Queue head should be 20")
	}
	stack.Dequeue()
	_, err := stack.Dequeue()
	if err == nil {
		t.Error("Empty queue should throw an error")
	}
}
