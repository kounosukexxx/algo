package l1011

import "sort"

// O(log(N*W) * W * log(N))
// 計算量は良さそうだけど、無視した部分が支配的になっている?

/*
 * @lc app=leetcode id=1011 lang=golang
 *
 * [1011] Capacity To Ship Packages Within D Days
 */

// @lc code=start
func shipWithinDays(weights []int, days int) int {
	sums := make([]int, 1, len(weights)+1)
	maxWeight := 0
	for i, w := range weights {
		sums = append(sums, sums[i]+w)
		maxWeight = max(maxWeight, w)
	}

	allSum := sums[len(sums)-1]
	if days == 1 {
		return allSum
	}

	l := max(maxWeight, (allSum+1)/days)
	r := allSum
	return sort.Search(r-l+1, func(i int) bool {
		return canShip(sums, days, i+l)
	}) + l
}

// O(days * log(|sums|)))
func canShip(sums []int, days, capacity int) bool {
	target := capacity
	for i := 0; i < days; i++ {
		lastShippedIndex := sort.Search(len(sums), func(i int) bool {
			return sums[i] > target
		}) - 1
		target = sums[lastShippedIndex] + capacity
	}

	return target == sums[len(sums)-1]+capacity
}

// これでやる場合は、comulative sumだとむずい。普通にやるべき
// func canShip2(sums []int, days, capacity int) bool {
// 	target := capacity
// 	passedDay := 0
// 	for i := 0; i < len(sums) && passedDay < days; i++ {
// 		sum := sums[i]
// 		if sum > target {
// 			passedDay++
// 			target = sums[i-1] + capacity
// 		}
// 		if i == len(sums)-1 {
// 			if sum < target {
// 				passedDay++
// 			}
// 		}
// 	}

// 	return passedDay <= days
// }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
