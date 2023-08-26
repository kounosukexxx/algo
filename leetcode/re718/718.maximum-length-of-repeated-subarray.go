package re718

// time: O(N*M)
// space: O(N)

/*
 * @lc app=leetcode id=718 lang=golang
 *
 * [718] Maximum Length of Repeated Subarray
 */

// @lc code=start
func findLength(nums1 []int, nums2 []int) int {
	subArrayCounts := make([]int, len(nums1))
	ans := 0
	for i := 0; i < len(nums1); i++ {
		tmpSubarrayCounts := make([]int, len(subArrayCounts))
		copy(tmpSubarrayCounts, subArrayCounts)
		for j := 0; j < len(nums2); j++ {
			if nums1[len(nums1)-1-i] == nums2[len(nums2)-1-j] {
				before := 0
				if j != 0 {
					before = tmpSubarrayCounts[j-1]
				}

				subArrayCounts[j] = before + 1
			} else {
				subArrayCounts[j] = 0
			}
		}

		ans = more(ans, max(subArrayCounts))
	}

	return ans
}

func more(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func max(a []int) int {
	res := 0
	for _, v := range a {
		res = more(res, v)
	}
	return res
}

// @lc code=end
