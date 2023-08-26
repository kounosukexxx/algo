package main

import "fmt"

// Queue implementation
type Queue struct {
	data []int
	size int
}

// NewQueue instantiates a new queue
func NewQueue(cap int) *Queue {
	return &Queue{data: make([]int, 0, cap), size: 0}
}

// Push adds a new element at the end of the queue
func (q *Queue) Push(n int) {
	q.data = append(q.data, n)
	q.size++
}

// Pop removes the first element from queue
func (q *Queue) Pop() bool {
	if q.IsEmpty() {
		return false
	}
	q.size--
	q.data = q.data[1:]
	return true
}

// Front returns the first element of queue
func (q *Queue) Front() int {
	return q.data[0]
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// String implements Stringer interface
func (q *Queue) String() string {
	return fmt.Sprint(q.data)
}
