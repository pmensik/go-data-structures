package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	Clear()
	for i := 1; i < 15; i++ {
		if !Add(i) {
			t.Error("Couldn't add elemento to the array")
		}
	}
	if capacity != 20 || len(array) != 20 {
		t.Error("Array capacity and length should be 20")
	}
}

func TestContains(t *testing.T) {
	Clear()
	Add(3)
	Add(33)
	if !Contains(33) || array[1] != 33 {
		t.Error("Array should contain element with value 33")
	}
	if Contains(4) {
		t.Error("Array should not contains element with value 4")
	}
}

func TestClear(t *testing.T) {
	Add(3)
	if !Contains(3) {
		t.Error("Array should contain element with value 3")
	}
	Clear()
	if Contains(3) {
		t.Error("Array should not contain element with value 3")
	}
}

func TestInsert(t *testing.T) {
	Clear()
	Add(3)
	Add(5)
	Add(8)
	Insert(20, 1)
	if array[1] != 20 {
		t.Error("Element with value should be under index 1")
	}
}

func TestRemove(t *testing.T) {
	Clear()
	Add(3)
	Add(5)
	Add(8)
	Add(15)
	Remove(1)
	if array[1] != 8 || array[2] != 15 {
		t.Error("Error removing element from index 1")
	}
}
