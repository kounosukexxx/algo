package l560

/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 */

// comurative sum: O(n^2)

// @lc code=start
func subarraySum(nums []int, k int) int {
	sums := make([]int, len(nums)+1)
	sums[0] = 0
	for i := 0; i < len(nums); i++ {
		sums[i+1] = sums[i] + nums[i]
	}

	ans := 0
	for i := 0; i < len(sums); i++ {
		for j := i + 1; j < len(sums); j++ {
			if sums[j]-sums[i] == k {
				ans++
			}
		}
	}

	return ans
}

// @lc code=end
