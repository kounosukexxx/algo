package l17

import (
	"math"
	"strconv"
)

/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start

func letterCombinations(digits string) []string {
	// digit-2 ➡︎ []string{...}
	var lettersIndexByDigit = [][]string{
		{
			"a", "b", "c",
		},
		{
			"d", "e", "f",
		},
		{
			"g", "h", "i",
		},
		{
			"j", "k", "l",
		},
		{
			"m", "n", "o",
		},
		{
			"p", "q", "r", "s",
		},
		{
			"t", "u", "v",
		},
		{
			"w", "x", "y", "z",
		},
	}
	ans := make(ansStr, 0, int(math.Pow(4, float64(len(digits)))))
	recurse(digits, &ans, "", lettersIndexByDigit)
	return ans
}

type ansStr []string

func recurse(digits string, ans *ansStr, curr string, letters [][]string) {
	if digits == "" {
		if curr != "" {
			*ans = append(*ans, curr)
		}
		return
	}

	head, _ := strconv.Atoi(string(digits[0]))
	letter := letters[head-2]
	for i := 0; i < len(letter); i++ {
		recurse(digits[1:], ans, curr+letter[i], letters)
	}
}

// @lc code=end
