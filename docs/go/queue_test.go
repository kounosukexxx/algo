package main

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue(10)
	fmt.Println(queue)

	if !queue.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", queue.IsEmpty())
	}

	queue.Push(10)
	fmt.Println(queue)
	queue.Push(1)
	fmt.Println(queue)
	queue.Push(-5)
	fmt.Println(queue)

	if queue.IsEmpty() {
		t.Fatalf("False is expected, but %v\n", queue.IsEmpty())
	}

	if queue.Front() != 10 {
		t.Fatalf("10 is expected, but %v\n", queue.Front())
	}

	queue.Pop()
	fmt.Println(queue)
	if queue.Front() != 1 {
		t.Fatalf("1 is expected, but %v\n", queue.Front())
	}

	queue.Pop()
	fmt.Println(queue)
	if queue.Front() != -5 {
		t.Fatalf("-5 is expected, but %v\n", queue.Front())
	}

	queue.Pop()
	fmt.Println(queue)
	if !queue.IsEmpty() {
		t.Fatalf("True is expected, but %v\n", queue.IsEmpty())
	}
}
