package l5

/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) string {
	indexesFromChar := make(map[rune][]int, len(s))
	for i, r := range s {
		indexesFromChar[r] = append(indexesFromChar[r], i)
	}
	l := 0
	r := 0
	maxLen := 0
	lWhenMax := 0
	rWhenMax := 0

	for _, indexes := range indexesFromChar {
		for i := range indexes {
			l = indexes[i]
			for j := len(indexes) - 1; j > i; j-- {
				r = indexes[j]
				if r-l+1 > maxLen && isPalindromic(s, l, r) {
					maxLen = r - l + 1
					lWhenMax = l
					rWhenMax = r
					continue
				}
			}
		}
	}

	return s[lWhenMax : rWhenMax+1]
}

func isPalindromic(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

// @lc code=end
