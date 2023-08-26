package heaptest

import (
	"container/heap"
	"fmt"
	"testing"
)

// NOTE: マジでh.Push, h.Popをそのまま使ってはいけない。heap.Push(h, x)ならまだ動く。
// h.push(), h.pop()を使え

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func Test_Heap(t *testing.T) {
	h := &intHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for !h.isEmpty() {
		fmt.Printf("%d ", heap.Pop(h))
	}

	h1 := newIntHeap([]int{2, 1, 5})
	fmt.Println(h1)
	h1.pushInt(3)
	fmt.Printf("minimum: %d\n", h1.front())
	for !h1.isEmpty() {
		fmt.Printf("%d ", h1.popInt())
	}

	h2 := newIntHeap(nil)
	h2.pushInt(30)
	h2.pushInt(15)
	h2.pushInt(20)
	fmt.Printf("minimum: %d\n", h2.front())
	for !h2.isEmpty() {
		fmt.Printf("%d ", h2.popInt())
	}
}
