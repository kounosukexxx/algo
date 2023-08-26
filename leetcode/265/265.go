package l265

import "math"

// we can optimize more
// O(N*k*k)
func minCostII(costs [][]int) int {
	currCosts := make([]int, len(costs[0]))
	for i, cost := range costs[0] {
		currCosts[i] = cost
	}

	for i := 1; i < len(costs); i++ {
		tmpCosts := make([]int, len(costs[i]))
		copy(tmpCosts, currCosts)
		for j := 0; j < len(costs[i]); j++ {
			currCosts[j] = min(tmpCosts, &j) + costs[i][j]
		}
	}

	return min(currCosts, nil)
}

func min(a []int, indexToExclude *int) int {
	min := math.MaxInt
	for i := 0; i < len(a); i++ {
		if indexToExclude != nil && i == *indexToExclude {
			continue
		}
		if a[i] < min {
			min = a[i]
		}
	}
	return min
}
