package queue

import (
	"errors"
)

type node struct {
	prev  *node
	value int
}

// Queue implementation based on linked list
type Queue struct {
	tail *node
}

// Enqueue puts element at the end of the queue
func (q *Queue) Enqueue(element int) bool {
	if q.IsEmpty() {
		q.tail = &node{prev: nil, value: element}
	} else {
		q.tail = &node{prev: q.tail, value: element}
	}
	return true
}

//Dequeue removes element from the front of the queue
func (q *Queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return -1, errors.New("Dequeue on empty queue not supported")
	}
	val := q.tail.value
	q.tail = q.tail.prev
	return val, nil
}

// Peek returns element from the front of the queue without removing it
func (q *Queue) Peek() int {
	return q.tail.value
}

// IsEmpty checks whether queue contains any elements
func (q *Queue) IsEmpty() bool {
	return q.tail == nil
}
