package l322

import "math"

// bottom up way

/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	dp := make([]int, 0, amount+1)
	dp = append(dp, 0)
	for i := 1; i < cap(dp); i++ {
		dp = append(dp, math.MaxInt)
	}

	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(coins); j++ {
			beforeAmount := i - coins[j]
			if beforeAmount < 0 || dp[beforeAmount] == math.MaxInt {
				continue
			}
			dp[i] = min(dp[i], dp[beforeAmount]+1)
		}
	}

	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
