package l209

/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	l := 0
	r := 0
	curr := 0
	ans := len(nums) + 1
	for r < len(nums) {
		curr += nums[r]
		for curr >= target {
			ans = min(ans, r-l+1)
			curr -= nums[l]
			l++
		}
		r++
	}
	if ans == len(nums)+1 {
		return 0
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
