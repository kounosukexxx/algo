package l121

/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	suffixMaxOfPrices := make([]int, len(prices))
	for i := len(prices) - 1; i >= 0; i-- {
		if i == len(prices)-1 {
			suffixMaxOfPrices[i] = prices[i]
		} else {
			suffixMaxOfPrices[i] = max(suffixMaxOfPrices[i+1], prices[i])
		}
	}

	maxProfit := 0
	for i, price := range prices {
		maxProfit = max(maxProfit, suffixMaxOfPrices[i]-price)
	}
	return maxProfit
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// @lc code=end
