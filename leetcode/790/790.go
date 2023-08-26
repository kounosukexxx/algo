package l790

const mod = 1e9 + 7

func numTilings(n int) int {
	if n == 1 {
		return 1
	}

	memo := []int{1, 1, 1, 1}
	preMemo := make([]int, 4)
	for i := 0; i < n-2; i++ {
		copy(preMemo, memo)
		memo[0] = (preMemo[0] + preMemo[3]) % mod
		memo[1] = memo[0]
		memo[2] = (preMemo[0] + preMemo[1] + preMemo[3]) % mod
		memo[3] = (preMemo[2] + preMemo[3]) % mod
	}

	return (memo[2] + memo[3]) % mod
}
