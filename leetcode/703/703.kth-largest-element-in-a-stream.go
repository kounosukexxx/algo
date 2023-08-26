package l703

import "sort"

// O(N^2)
/*
 * @lc app=leetcode id=703 lang=golang
 *
 * [703] Kth Largest Element in a Stream
 */

// @lc code=start
type KthLargest struct {
	data []int
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	return KthLargest{
		data: nums,
		k:    k,
	}
}

func (this *KthLargest) Add(val int) int {
	i := sort.SearchInts(this.data, val)

	// ダメな例
	// suffix := this.data[i:]
	// this.data = append(this.data[:i], val)
	// this.data = append(this.data, suffix...)

	suffix := make([]int, len(this.data[i:]))
	copy(suffix, this.data[i:])
	this.data = append(this.data[:i], val)
	this.data = append(this.data, suffix...)

	// 1,2,3,4,5
	return this.data[len(this.data)-this.k]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
// @lc code=end
