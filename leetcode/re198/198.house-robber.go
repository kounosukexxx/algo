package re198

/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
func rob(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = -1
	}

	if len(nums) == 1 {
		return nums[0]
	}

	setDp(dp, len(nums)-1, nums)
	setDp(dp, len(nums)-2, nums)
	return max(dp[len(nums)-1], dp[len(nums)-2])
}

// i should be >= 0
func setDp(dp []int, i int, nums []int) {
	if dp[i] >= 0 {
		return
	}

	switch {
	case i == 0, i == 1:
		dp[i] = nums[i]
	case i == 2:
		dp[i] = nums[0] + nums[2]
	default:
		setDp(dp, i-2, nums)
		setDp(dp, i-3, nums)
		dp[i] = nums[i] + max(dp[i-2], dp[i-3])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
