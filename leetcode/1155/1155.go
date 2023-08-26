package l1155

const mod = 1e9 + 7

func numRollsToTarget(n int, k int, target int) int {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, target)
		for j := 0; j < target; j++ {
			dp[i][j] = -1
		}
	}

	setDP(dp, n-1, target-1, k)
	return dp[n-1][target-1]
}

func setDP(dp [][]int, i, j, K int) {
	if dp[i][j] > -1 {
		// already visited
		return
	}

	if i == 0 {
		if j < K {
			dp[i][j] = 1
			return
		}

		dp[i][j] = 0
		return
	}

	val := 0
	for k := 1; k <= K; k++ {
		if j-k < 0 {
			break
		}

		setDP(dp, i-1, j-k, K)
		val += dp[i-1][j-k]
		val %= mod
	}

	dp[i][j] = val
}
