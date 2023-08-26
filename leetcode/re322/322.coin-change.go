package re322

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	sort.Slice(coins, func(i, j int) bool { return coins[i] < coins[j] })
	dp := make([]int, 0, amount+1)
	dp = append(dp, 0)
	for i := 1; i < cap(dp); i++ {
		dp = append(dp, math.MaxInt)
	}
	setDp(dp, coins, amount)
	if dp[amount] == math.MaxInt-1 { // unreachable
		return -1
	}
	return dp[amount]
}

// dp: init value: math.MaxInt, non reachable: math.MaxInt-1
func setDp(dp []int, coins []int, amount int) {
	if dp[amount] == math.MaxInt-1 || dp[amount] != math.MaxInt {
		return
	}

	for _, coin := range coins {
		beforeAmount := amount - coin
		if beforeAmount < 0 {
			break // coinsはsort済みなので、continueではなくbreak
		}
		setDp(dp, coins, beforeAmount)
		if dp[beforeAmount] == math.MaxInt || dp[beforeAmount] == math.MaxInt-1 {
			continue
		}
		dp[amount] = min(dp[amount], dp[beforeAmount]+1)
	}

	if dp[amount] == math.MaxInt {
		dp[amount] = math.MaxInt - 1 // 到達不可能としてmark
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
