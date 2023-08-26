package l1473

import "math"

/*
 * @lc app=leetcode id=1473 lang=golang
 *
 * [1473] Paint House III
 */

// 添字じゅんばん間違えた。
// @lc code=start
func minCost(houses []int, cost [][]int, m int, n int, target int) int {
	dp := make([][][]int, m)
	for i := range houses {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, target)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	for j := 0; j < n; j++ {
		dp[0][j][0] = cost[0][j]
	}

	for i := 0; i < m-1; i++ {
		for j := 0; j < n; j++ {
			c := cost[i][j]
			for k := 0; k < target; k++ {
				if dp[i][j][k] == -1 {
					continue // skip because unreachable
				}

				dp[i+1][j][k] = min(dp[i+1][j][k], dp[i][j][k]+c)

				if k == target-1 {
					continue
				}

				dp[i+1][j][k+1] = minWithExcludeIndex(dp[i], j, k) + c
			}
		}
	}

	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minWithExcludeIndex(a [][]int, exIndex, k int) int {
	m := math.MaxInt
	for j := range a {
		if j == exIndex {
			continue
		}

		m = min(m, a[j][k])
	}

	return m
}

// @lc code=end
