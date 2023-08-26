package l15

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	prev := math.MaxInt
	ans := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		curr := nums[i]
		if curr == prev {
			continue
		}
		prev = curr

		sums := twoSum(nums[i+1:], -curr)
		for _, sum := range sums {
			ans = append(ans, []int{curr, sum[0], sum[1]})
		}
	}

	return ans
}

func twoSum(nums []int, target int) [][]int {
	l := 0
	r := len(nums) - 1
	prevAccptedLeftVal := math.MaxInt
	res := [][]int{}
	for l < r {
		leftVal := nums[l]
		if prevAccptedLeftVal == leftVal {
			l++
			continue
		}

		sum := leftVal + nums[r]
		if sum == target {
			res = append(res, []int{leftVal, nums[r]})
			l++
			r--
			prevAccptedLeftVal = leftVal
		} else if sum < target {
			l++
		} else {
			r--
		}
	}

	return res
}

// @lc code=end
