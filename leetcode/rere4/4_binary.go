package l4

import "math"

/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	isNumsEven := (len(nums1)+len(nums2))%2 == 0
	mid := (len(nums1) + len(nums2) + 1) / 2

	A := nums1 // the len is smaller
	B := nums2
	if len(nums1) > len(nums2) {
		A, B = B, A
	}

	l := 0
	r := len(A) - 1
	if r == -1 {
		l = -1 // we want a to become -1
	}

	for {
		a := (l + r) / 2
		if r < 0 {
			a = -1 // edge case
		}
		b := mid - a - 2
		var leftA, rightA, leftB, rightB int

		if a < 0 {
			leftA = math.MinInt
		} else {
			leftA = A[a]
		}
		if a+1 > len(A)-1 {
			rightA = math.MaxInt
		} else {
			rightA = A[a+1]
		}

		if b < 0 {
			leftB = math.MinInt
		} else {
			leftB = B[b]
		}
		if b+1 > len(B)-1 {
			rightB = math.MaxInt
		} else {
			rightB = B[b+1]
		}

		// dividion succeeded
		if leftA <= rightB && leftB <= rightA {
			if isNumsEven {
				return float64(max(leftA, leftB)+min(rightA, rightB)) / 2
			}
			return float64(max(leftA, leftB))
		}

		if leftB > rightA {
			l = a + 1 // go right
		} else {
			r = a - 1 // go left
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
