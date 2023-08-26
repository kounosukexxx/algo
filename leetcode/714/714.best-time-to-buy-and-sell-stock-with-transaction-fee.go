package l714

/*
 * @lc app=leetcode id=714 lang=golang
 *
 * [714] Best Time to Buy and Sell Stock with Transaction Fee
 */

// @lc code=start
// space O(1)
func maxProfit(prices []int, fee int) int {
	length := len(prices)
	maxMoney0 := 0
	maxMoney1 := -prices[0] - fee
	for i := 1; i < length; i++ {
		yesterdayMoney0 := maxMoney0
		yesterdayMoney1 := maxMoney1
		maxMoney0 = max(yesterdayMoney0, yesterdayMoney1+prices[i])
		maxMoney1 = max(yesterdayMoney1, yesterdayMoney0-prices[i]-fee)
	}

	return max(maxMoney0, maxMoney1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// array. space O(N)
// func maxProfit(prices []int, fee int) int {
// 	length := len(prices)
// 	dp := make([][]int, 0, length)
// 	dp = append(dp, []int{0, -prices[0] - fee})
// 	for i := 1; i < length; i++ {
// 		dp = append(dp, []int{
// 			max(dp[i-1][0], dp[i-1][1]+prices[i]),
// 			max(dp[i-1][1], dp[i-1][0]-prices[i]-fee),
// 		})
// 	}

// 	return max(dp[length-1][0], dp[length-1][0])
// }

// @lc code=end
