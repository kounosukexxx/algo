package l373

import "math"

// ダメ

/*
 * @lc app=leetcode id=373 lang=golang
 *
 * [373] Find K Pairs with Smallest Sums
 */

// @lc code=start
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	even := 0
	odd := 1
	a := 0
	b := 0

	ans := make([][]int, 0)

	for i := 0; i < k; i++ {
		sum1 := getSum(even, a, nums1, nums2)
		sum2 := getSum(odd, b, nums1, nums2)
		if sum1 == math.MaxInt && sum2 == math.MaxInt {
			break
		}

		if sum1 <= sum2 {
			ans = append(ans, []int{nums1[even], nums2[a]})
			transferIndex(&even, &a, nums2)
		} else {
			ans = append(ans, []int{nums1[odd], nums2[b]})
			transferIndex(&odd, &b, nums2)
		}
	}

	return ans
}

func getSum(index1, index2 int, nums1, nums2 []int) int {
	if index1 >= len(nums1) {
		return math.MaxInt
	}
	return nums1[index1] + nums2[index2]
}

func transferIndex(index1, index2 *int, nums2 []int) {
	if *index2+1 == len(nums2) {
		*index1 += 2
		*index2 = 0
	} else {
		*index2++
	}
}

// @lc code=end
