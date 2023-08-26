package l139

import "strings"

/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// dp
// @lc code=start
func wordBreak(s string, wordDict []string) bool {
	checkedSuffixLength := make([]bool, len(s))
	return internal(checkedSuffixLength, s, wordDict)
}

func internal(checkedSuffixLength []bool, s string, wordDict []string) bool {
	if checkedSuffixLength[len(s)-1] {
		return false
	}
	checkedSuffixLength[len(s)-1] = true

	for _, word := range wordDict {
		if word == s {
			return true
		}
		if strings.HasPrefix(s, word) {
			if internal(checkedSuffixLength, s[len(word):], wordDict) {
				return true
			}
		}
	}
	return false
}

// @lc code=end
