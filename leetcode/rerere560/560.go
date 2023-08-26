package rere560

/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 */

// @lc code=start

// loopをもっと再利用できる。

// editorial4

func subarraySum(nums []int, k int) int {
	sums := make([]int, len(nums)+1)
	mp := make(map[int]int, len(sums))
	ans := 0

	sums[0] = 0
	mp[0] = 1
	for i := 0; i < len(nums); i++ {
		sum := sums[i] + nums[i]
		sums[i+1] = sum

		ans += mp[sum-k]

		mp[sum]++
	}

	return ans
}

// @lc code=end
