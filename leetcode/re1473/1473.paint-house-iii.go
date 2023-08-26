package re1473

import "math"

/*
 * @lc app=leetcode id=1473 lang=golang
 *
 * [1473] Paint House III
 */

// @lc code=start
// もうちょっとのとこであきらめる
func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	dp := make([][][]int, m)
	for i := range houses {
		dp[i] = make([][]int, target)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
			for k := range dp[i][j] {
				dp[i][j][k] = math.MaxInt
			}
		}
	}

	// TODO:
	for i, c := range cost[0] {
		dp[0][0][i] = c
	}

	for i := 0; i < m-1; i++ {
		paintedColor := houses[i+1] - 1
		isAleacyPainted := paintedColor != -1
		for j := 0; j < target; j++ {
			for k := 0; k < n; k++ {
				// if dp[i][j][k] == math.MaxInt {
				// 	continue // skip because unreachable
				// }

				if isAleacyPainted {
					if paintedColor != k {
						continue // k色にはぬれない
					}

				}

				c := cost[i+1][k]
				if isAleacyPainted {
					c = 0
				}

				// 直前と同じ色を塗る場合
				if dp[i][j][k] != math.MaxInt {
					// paint if not unreachable
					dp[i+1][j][k] = min(dp[i+1][j][k], dp[i][j][k]+c)
				}

				// 直前と違う色を塗る場合
				if j+1 <= target-1 {
					m := minWithExcludeIndex(dp[i][j], &k)
					if m == math.MaxInt {
						continue // skip because unreachable
					}
					dp[i+1][j+1][k] = m + c
				}
			}
		}
	}

	ans := minWithExcludeIndex(dp[m-1][target-1], nil)
	if ans == math.MaxInt {
		ans = -1
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minWithExcludeIndex(a []int, exIndex *int) int {
	ans := math.MaxInt
	for i, v := range a {
		if exIndex != nil && i == *exIndex {
			continue
		}
		ans = min(ans, v)
	}
	return ans
}

// @lc code=end
