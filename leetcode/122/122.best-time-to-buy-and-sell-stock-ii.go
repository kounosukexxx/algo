package l122

/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// NOTE: 貪欲法
// prices[i] == prices[i-1] かつ最後の時、
// 谷を越えた麓ならば、profitを加算する必要はない。
// 上がっていたときのみprofitを加算する。
// maxPriceをした意味は今回特にない。0でもいい。

// re: while (i < len-1)で、
// ①一回下がり続ける。②その後上がり続ける。（どっちもwhile文）

// @lc code=start
func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxPrice := -1
	upInPreTurn := false
	profit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] == prices[i-1] {
			if i == len(prices)-1 && upInPreTurn {
				profit += maxPrice - minPrice
			}
			continue
		}

		up := prices[i] > prices[i-1]
		if up {
			maxPrice = prices[i]
			upInPreTurn = true
			if i == len(prices)-1 {
				profit += maxPrice - minPrice
			}
		} else {
			if upInPreTurn {
				profit += maxPrice - minPrice
				maxPrice = -1 // reset
			}
			minPrice = prices[i]
			upInPreTurn = false
		}

	}

	return profit
}

// @lc code=end
