package l213

/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 */

// re: 無駄なくせる

// -2: unset
// -1: invalid
// >=0: valid
// @lc code=start
func rob(nums []int) int {
	switch len(nums) {
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	}

	dp1 := newDp(nums)
	dp2 := newDp(nums)
	dp3 := newDp(nums)

	setDp1(dp1, len(nums)-2, nums)
	setDp1(dp1, len(nums)-3, nums)
	ans1 := max(dp1[len(nums)-2], dp1[len(nums)-3])

	setDp2(dp2, len(nums)-1, nums)
	setDp2(dp2, len(nums)-2, nums)
	ans2 := max(dp2[len(nums)-1], dp2[len(nums)-2])

	setDp3(dp3, len(nums)-1, nums)
	setDp3(dp3, len(nums)-2, nums)
	ans3 := max(dp3[len(nums)-1], dp3[len(nums)-2])

	return max(ans1, max(ans2, ans3))
}

func newDp(nums []int) []int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = -2
	}
	return dp
}

func setDp1(dp []int, i int, nums []int) {
	if dp[i] >= -1 {
		return
	}

	switch i {
	case 0:
		dp[i] = nums[i]
	case 1:
		dp[i] = -1
	case 2:
		dp[i] = nums[0] + nums[2]
	default:
		setDp1(dp, i-2, nums)
		setDp1(dp, i-3, nums)
		dp[i] = nums[i] + max(dp[i-2], dp[i-3])
	}
}

func setDp2(dp []int, i int, nums []int) {
	if dp[i] >= -1 {
		return
	}

	switch i {
	case 0:
		dp[i] = -1
	case 1:
		dp[i] = nums[i]
	case 2:
		dp[i] = -1
	default:
		setDp2(dp, i-2, nums)
		setDp2(dp, i-3, nums)
		dp[i] = nums[i] + max(dp[i-2], dp[i-3])
	}
}

func setDp3(dp []int, i int, nums []int) {
	if dp[i] >= -1 {
		return
	}

	switch i {
	case 0, 1, 3:
		dp[i] = -1
	case 2:
		dp[i] = nums[i]
	default:
		setDp3(dp, i-2, nums)
		setDp3(dp, i-3, nums)
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
