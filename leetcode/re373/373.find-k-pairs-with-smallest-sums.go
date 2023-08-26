package re373

import "container/heap"

/*
 * @lc app=leetcode id=373 lang=golang
 *
 * [373] Find K Pairs with Smallest Sums
 */

// だめ。数学的に正しくない。
// @lc code=start
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	heap := &intHeap{}
	ans := make([][]int, 0)
	maxIndex1 := len(nums1) - 1
	maxIndex2 := len(nums2) - 1

	for i := 0; ; i++ {
		index1 := 0
		index2 := i
		if index2 > maxIndex2 {
			index2 = maxIndex2
			maxIndex1 = index2 - maxIndex2
		}
		for index1 <= i && index1 <= maxIndex1 {
			heap.pushInt([]int{nums1[index1], nums2[index2]})
			index1++
			index2--
		}

		for heap.len() > 0 {
			ans = append(ans, heap.popInt())
			if len(ans) == k {
				return ans
			}
		}
	}
}

// An intHeap is a min-heap of 2 ints.
// size of intHeap is defined by sum of 2 ints.
type intHeap [][]int

func (h *intHeap) pushInt(x []int) {
	heap.Push(h, x)
}

func (h *intHeap) popInt() []int {
	return heap.Pop(h).([]int)
}

func (h *intHeap) front() []int {
	heap := *h
	return heap[0]
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
func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i][0]+h[i][1] < h[j][0]+h[j][1] } // for min-heap
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// @lc code=end
