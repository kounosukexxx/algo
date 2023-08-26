package l4

import "sort"

/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start

// 複雑すぎてむりやった
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	isNumsEven := (len(nums1)+len(nums2))%2 == 0
	if len(nums1) == 0 {
		return getAns(nums1, nums2, border1, border2)
	}
	m := (len(nums1) + len(nums2) - 1) / 2

	border1 := (len(nums1) - 1) / 2

	sort.Search()
	for {
		border2 := m - border1 - 1
		switch {
		case border1 == len(nums1)-1:
			return getAns(nums1, nums2, border1, border2)
		case border2+1 < len(nums2) && nums1[border1] > nums2[border2+1]:
			border1--
		case nums2[border2] > nums1[border1+1]:
			border1++
		default:
			return getAns(nums1, nums2, border1, border2)
		}
	}
}

func getAns(nums1, nums2 []int, border1, border2 int, isNumsEven bool) float64 {
	if border1 == len(nums1)-1 {
		if isNumsEven {
			return float64(nums1[border1]+nums2[0]) / 2
		}
		return float64(nums1[border1])
	}
	if border1 < 0 {
		if isNumsEven {
			return float64(nums2[border2]+nums1[0]) / 2
		}
	}

	if isNumsEven {
		return (float64(max(nums1[border1], nums2[border2])) +
			float64(min(nums1[border1+1], nums2[border2+1]))) / 2
	}
	return float64(max(nums1[border1], nums2[border2]))
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
