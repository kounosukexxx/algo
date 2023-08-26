package l256

func minCost(costs [][]int) int {
	redCost := costs[0][0]
	blueCost := costs[0][1]
	greenCost := costs[0][2]
	for i := 1; i < len(costs); i++ {
		tmpRedCost := redCost
		tmpBlueCost := blueCost
		tmpGreenCost := greenCost
		redCost = min(tmpBlueCost, tmpGreenCost) + costs[i][0]
		blueCost = min(tmpRedCost, tmpGreenCost) + costs[i][1]
		greenCost = min(tmpRedCost, tmpBlueCost) + costs[i][2]
	}

	return min(min(redCost, blueCost), greenCost)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
