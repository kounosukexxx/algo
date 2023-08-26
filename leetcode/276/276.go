package l276

func numWays(n int, k int) int {
	if n == 1 {
		return k
	}
	dp0 := make([]int, n)
	dp1 := make([]int, n)
	dp0[1] = k * (k - 1)
	dp1[1] = k
	for i := 1; i < n-1; i++ {
		dp0[i+1] = dp0[i]*(k-1) + dp1[i]*(k-1)
		dp1[i+1] = dp0[i]
	}
	return dp0[n-1] + dp1[n-1]
}
