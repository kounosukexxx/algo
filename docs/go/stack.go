package main

import "fmt"

// Stack implementation
type Stack struct {
	data []int
	size int
}

// NewStack instantiates a new stack
func NewStack(cap int) *Stack {
	return &Stack{data: make([]int, 0, cap), size: 0}
}

// Push adds a new element at the end of the stack
func (s *Stack) Push(n int) {
	s.data = append(s.data, n)
	s.size++
}

// Pop removes the last element from stack
func (s *Stack) Pop() bool {
	if s.IsEmpty() {
		return false
	}
	s.size--
	s.data = s.data[:s.size]
	return true
}

// Top returns the last element of stack
func (s *Stack) Top() int {
	return s.data[s.size-1]
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

// String implements Stringer interface
func (s *Stack) String() string {
	return fmt.Sprint(s.data)
}
