package rere373

import "container/heap"

/*
 * @lc app=leetcode id=373 lang=golang
 *
 * [373] Find K Pairs with Smallest Sums
 */

// @lc code=start
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	heap := initIntHeap(nums1, nums2)
	heap.pushInt([]int{0, 0})

	ans := make([][]int, 0, k)
	i := 0
	for !heap.isEmpty() {
		popped := heap.popInt()
		ans = append(ans, []int{nums1[popped[0]], nums2[popped[1]]})
		if len(ans) == k {
			break
		}

		if popped[1]+1 < len(nums2) {
			heap.pushInt([]int{popped[0], popped[1] + 1})
		}

		if popped[0]+popped[1] == i {
			i++
			if i <= len(nums1)-1 {
				heap.pushInt([]int{i, 0})
			}
		}
	}

	return ans
}

// An intHeap is a min-heap of 2 ints.
// size of intHeap is defined by sum of 2 ints.
type intHeap struct {
	data  [][]int
	nums1 []int
	nums2 []int
}

func initIntHeap(nums1, nums2 []int) *intHeap {
	return &intHeap{
		nums1: nums1,
		nums2: nums2,
	}
}

func (h *intHeap) pushInt(x []int) {
	heap.Push(h, x)
}

func (h *intHeap) popInt() []int {
	return heap.Pop(h).([]int)
}

func (h *intHeap) front() []int {
	heap := *h
	return heap.data[0]
}

func (h *intHeap) isEmpty() bool {
	return h.Len() == 0
}

func (h *intHeap) len() int {
	return h.Len()
}

// implement heap.Interface
// DON'T use below methods directly
// these methods are used inside container/heap package
func (h intHeap) Len() int { return len(h.data) }
func (h intHeap) Less(i, j int) bool {
	return h.nums1[h.data[i][0]]+h.nums2[h.data[i][1]] < h.nums1[h.data[j][0]]+h.nums2[h.data[j][1]]
}                               // for min-heap
func (h intHeap) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *intHeap) Push(x interface{}) {
	h.data = append(h.data, x.([]int))
}
func (h *intHeap) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

// @lc code=end
