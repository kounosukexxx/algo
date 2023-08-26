package l53

import "math"

func maxSubArray(nums []int) int {
	curr := 0
	res := math.MinInt32

	for _, num := range nums {
		curr = max(curr+num, num)
		res = max(res, curr)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
