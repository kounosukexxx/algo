package l253

import (
	"container/heap"
	"sort"
)

// editorialちゃんと読んで無いけど、多分同じ
// goでのheapパッケージの使い方注意。自分で定義したprivate methodを使用すること。

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	h := newIntHeap(nil)

	ans := 0
	for i := 0; i < len(intervals); i++ {
		for !h.isEmpty() && intervals[i][0] >= h.front() {
			h.popInt()
		}
		h.pushInt(intervals[i][1])
		ans = max(ans, h.len())
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

func (h *intHeap) isEmpty() bool {
	return h.Len() == 0
}

func (h intHeap) len() int {
	return h.Len()
}

// implement heap.Interface
// DON'T use below methods directly
func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] } // for min-heap
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
