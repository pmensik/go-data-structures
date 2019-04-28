package linkedlist

type node struct {
	next, prev *node
	value      int
}

// DoubleLinkedList represents double linked list data structure
type DoubleLinkedList struct {
	head, tail *node
	size       int
}

// Empty returns whether list is empty
func (list *DoubleLinkedList) Empty() bool {
	return list.size == 0
}

// AddFirst will insert element at the beginning of the list
func (list *DoubleLinkedList) AddFirst(element int) bool {
	if list.Empty() {
		n := &node{next: nil, prev: nil, value: element}
		list.head, list.tail = n, n
	} else {
		list.head.prev = &node{next: nil, prev: list.head, value: element}
		list.head = list.head.prev
	}
	list.size++
	return true
}

// AddLast will insert element to the end of the list
func (list *DoubleLinkedList) AddLast(element int) bool {
	if list.Empty() {
		n := &node{next: nil, prev: nil, value: element}
		list.head, list.tail = n, n
	} else {
		list.tail.next = &node{next: nil, prev: list.tail, value: element}
		list.tail = list.tail.next
	}
	list.size++
	return true
}

// Add will insert element at tail of the list
func (list *DoubleLinkedList) Add(element int) bool {
	return list.AddLast(element)
}

// Clear remove all elements from the list
func (list *DoubleLinkedList) Clear() {
	n := list.head
	for n != nil {
		next := n.next
		n.prev, n.next = nil, nil
		n.value = 0
		n = next
	}
	list.head, list.tail = nil, nil
	list.size = 0
}

// Contains checks whether element is present in list
func (list *DoubleLinkedList) Contains(element int) bool {
	n := list.head
	for n != nil {
		if n.value == element {
			return true
		}
		n = n.next
	}
	return false
}

// RemoveAt removes element from a position in the list
func (list *DoubleLinkedList) RemoveAt(index int) int {
	n := list.head
	for i := 0; i < index; i++ {
		n = n.next
	}
	n.next.prev = n.prev
	n.prev.next = n.next
	return n.value
}

// Get returns an element from a position in the list
func (list *DoubleLinkedList) Get(index int) int {
	n := list.head
	for i := 0; i < index; i++ {
		n = n.next
	}
	return n.value
}
