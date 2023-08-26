package l45

/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}

	curr := 0
	jumpCount := 0
	for {
		jumpCount++

		// when you can reach goal after a jump
		if curr+nums[curr] >= len(nums)-1 {
			return jumpCount
		}

		// when you can't reach goal after a jump
		curr += toJump(nums, curr, nums[curr])
	}
}

func toJump(nums []int, from, jumpCap int) int {
	max := -1
	var toJump int
	for i := 1; i <= jumpCap; i++ {
		num := nums[from+i] + i
		if num > max {
			max = num
			toJump = i
		}
	}

	return toJump
}

// @lc code=end
