package l779

/*
 * @lc app=leetcode id=779 lang=golang
 *
 * [779] K-th Symbol in Grammar
 */

// @lc code=start
func kthGrammar(n int, k int) int {
	return getBit(k, 0)
}

func getBit(index, val int) int {
	if index == 1 {
		if val == 0 {
			return 0
		}
		return 1
	}

	parentVal := 0
	if (index%2 == 0 && val == 0) || (index%2 == 1 && val == 1) {
		parentVal = 1
	}
	parentIndex := index / 2
	if index%2 == 1 {
		parentIndex++
	}
	return getBit(parentIndex, parentVal)
}

// @lc code=end
