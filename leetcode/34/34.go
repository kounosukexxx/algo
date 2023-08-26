package l34

import "sort"

/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	startIndex := sort.Search(len(nums), func(i int) bool {
		return nums[i] >= target
	})
	if startIndex == len(nums) {
		return []int{-1, -1}
	}
	if nums[startIndex] != target {
		return []int{-1, -1}
	}
	endIndex := reverseSearch(len(nums), func(i int) bool {
		return nums[i] <= target
	})

	return []int{startIndex, endIndex}
}

// fを満たす最大のiを返す
// 一つも見つからなければ、-1を返す
func reverseSearch(n int, f func(int) bool) int {
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1)
		if uint(i+j)%2 == 1 {
			h++
		}

		if h == n || !f(h) {
			j = h - 1
		} else {
			i = h
		}
	}

	if i == 0 && !f(0) {
		return -1
	}
	return i
}

// @lc code=end
