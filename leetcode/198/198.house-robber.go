package l198

/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// timeout and policy is incorrect

// @lc code=start
func rob(nums []int) int {
	var ans int
	recurse(&ans, 0, nums)
	return ans
}

func recurse(ans *int, curr int, nums []int) {
	if len(nums) == 0 {
		if curr > *ans {
			*ans = curr
		}
		return
	}
	if len(nums) == 1 {
		if curr+nums[0] > *ans {
			*ans = curr + nums[0]
		}
		return
	}
	if len(nums) == 2 {
		toAdd := nums[0]
		if nums[1] > toAdd {
			toAdd = nums[1]
		}

		if curr+toAdd > *ans {
			*ans = curr + toAdd
		}
		return
	}
	if len(nums) > 2 {
		recurse(ans, curr+nums[1], nums[3:])
	}
	recurse(ans, curr+nums[0], nums[2:])
}

// @lc code=end
