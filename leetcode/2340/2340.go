package l2340

import "math"

func minimumSwaps(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	minVal := math.MaxInt32
	var minIndex int
	maxVal := math.MinInt32
	var maxIndex int
	for i, num := range nums {
		if num < minVal {
			minVal = num
			minIndex = i
		}
		if num >= maxVal {
			maxVal = num
			maxIndex = i
		}
	}

	// all numbers are the same
	if minVal == maxVal {
		return 0
	}

	res := len(nums) - 1 - maxIndex + minIndex
	if minIndex > maxIndex {
		return res - 1
	}
	return res
}
