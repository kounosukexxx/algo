package re703

import "container/heap"

// O(N^2)
/*
 * @lc app=leetcode id=703 lang=golang
 *
 * [703] Kth Largest Element in a Stream
 */

// @lc code=start
type KthLargest struct {
	heap *intHeap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	heap := newIntHeap(nums)
	return KthLargest{
		heap: heap,
		k:    k,
	}
}

func (kl *KthLargest) Add(val int) int {
	kl.heap.pushInt(val)
	for kl.heap.len() > kl.k {
		kl.heap.popInt()
	}
	return kl.heap.front()
}

// An intHeap is a min-heap of ints.
type intHeap []int

func newIntHeap(array []int) *intHeap {
	h := make(intHeap, 0)
	for _, v := range array {
		h = append(h, v)
	}
	heap.Init(&h)
	return &h
}

func (h *intHeap) pushInt(x int) {
	heap.Push(h, x)
}

func (h *intHeap) popInt() int {
	return heap.Pop(h).(int)
}

func (h *intHeap) front() int {
	heap := *h
	return heap[0]
}

func (h *intHeap) len() int {
	return h.Len()
}

// implement heap.Interface
// DON'T use below methods directly
// these methods are used inside container/heap package
func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] } // for min-heap
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
// @lc code=end
