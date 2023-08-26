package re1011

import "sort"

/*
 * @lc app=leetcode id=1011 lang=golang
 *
 * [1011] Capacity To Ship Packages Within D Days
 */

// @lc code=start
func shipWithinDays(weights []int, days int) int {
	maxWeight := 0
	totalLoad := 0
	for _, w := range weights {
		totalLoad += w
		maxWeight = max(maxWeight, w)
	}

	if days == 1 {
		return totalLoad
	}

	l := max(maxWeight, (totalLoad+1)/days)
	r := totalLoad
	return sort.Search(r-l+1, func(i int) bool {
		return canShip(weights, days, i+l)
	}) + l
}

// O(N)
func canShip(weights []int, days int, capacity int) bool {
	daysNeeded := 1
	currLoad := 0
	for _, w := range weights {
		currLoad += w
		if currLoad > capacity {
			currLoad = w
			daysNeeded++
			if daysNeeded > days {
				return false
			}
		}
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
