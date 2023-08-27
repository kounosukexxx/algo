package l1143

func longestCommonSubsequence(text1 string, text2 string) int {
	memo := make([][]int, 0, len(text1))
	for i := 0; i < len(text1); i++ {
		memo = append(memo, make([]int, len(text2)))
	}

	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			memo[i][j] = -1
		}
	}

	return recurse(text1, text2, 0, 0, memo)
}

func recurse(text1, text2 string, i, j int, memo [][]int) int {
	if i >= len(text1) || j >= len(text2) {
		return 0
	}

	m := memo[i][j]
	if m != -1 {
		return m
	}

	res := 0
	// when use r1
	for i2, r2 := range text2[j:] {
		if text1[i] == byte(r2) {
			res = max(res, 1+recurse(text1, text2, i+1, j+i2+1, memo))
			break
		}
	}
	// when not use r1
	res = max(res, recurse(text1, text2, i+1, j, memo))
	// memo
	memo[i][j] = res

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
