package l11

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	ans := 0
	for l < r {
		ans = max(ans, getArea(height, l, r))
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getArea(height []int, l, r int) int {
	return (r - l) * min(height[l], height[r])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
