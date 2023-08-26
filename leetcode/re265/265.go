package re265

import "math"

// 計算量的に良くなっているけど、遅くなってる
// O(N*K + K)
func minCostII(costs [][]int) int {
	minCost, secondMinCost, minColor := getMins(costs[0])

	for i := 1; i < len(costs); i++ {
		currCosts := make([]int, 0, len(costs[i]))

		for j := 0; j < len(costs[i]); j++ {
			toAdd := minCost
			if j == minColor {
				toAdd = secondMinCost
			}
			currCosts = append(currCosts, toAdd+costs[i][j])
		}

		minCost, secondMinCost, minColor = getMins(currCosts)
	}

	return minCost
}

func getMins(a []int) (int, int, int) {
	min := math.MaxInt
	secondMin := math.MaxInt
	minIndex := 0
	for i, v := range a {
		if v < min {
			secondMin = min
			min = v
			minIndex = i
		} else if v < secondMin {
			secondMin = v
		}
	}

	return min, secondMin, minIndex
}
