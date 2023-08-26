package l300

/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	len := len(nums)
	dp := make([]int, len)
	for i := 0; i < len; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[len-1-j] > nums[len-1-i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	ans := dp[0]
	for i := 1; i < len; i++ {
		ans = max(ans, dp[i])
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
