package linkedlist

import (
	"testing"
)

func TestAddFirst(t *testing.T) {
	list := &DoubleLinkedList{}
	list.AddFirst(20)
	if list.size != 1 {
		t.Error("List size should be 1 after adding first element")
	}
	list.AddFirst(30)
	if list.head.value != 30 {
		t.Error("Value of head should be 30")
	}
}

func TestAdd(t *testing.T) {
	list := &DoubleLinkedList{}
	list.Add(20)
	list.Add(30)
	list.Add(40)
	list.Add(50)
	if list.size != 4 {
		t.Error("List size should be 4 after adding some elements")
	}
	if list.tail.value != 50 {
		t.Error("Last element should be 50")
	}
	if list.head.value != 20 {
		t.Error("First element should be 20")
	}
}

func TestClear(t *testing.T) {
	list := &DoubleLinkedList{}
	list.Add(20)
	list.Add(30)
	list.Add(40)
	list.Add(50)
	list.Clear()
	if list.size != 0 {
		t.Error("List size should be 0 after cleaning")
	}
	if list.head != nil {
		t.Error("First element should be nil")
	}
	if list.tail != nil {
		t.Error("Last element should be nil")
	}
}

func TestContains(t *testing.T) {
	list := &DoubleLinkedList{}
	list.Add(3)
	list.Add(33)
	if !list.Contains(33) {
		t.Error("List should contain element with value 33")
	}
	if list.Contains(4) {
		t.Error("List should not contain element with value 4")
	}
}

func TestRemoveAt(t *testing.T) {
	list := &DoubleLinkedList{}
	list.Add(3)
	list.Add(33)
	list.Add(52)
	list.Add(21)
	list.RemoveAt(2)
	if !(list.Get(2) == 21) {
		t.Error("List should contain element with value 21 after removal")
	}
}
