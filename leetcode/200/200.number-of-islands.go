package l200

/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	landed := newLanded(grid)
	ans := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				if !landed[i][j] {
					ans++
					landed.markLandAsLanded(i, j, grid)
				}
			}
		}
	}
	return ans
}

type landed [][]bool

func newLanded(grid [][]byte) landed {
	l := make(landed, len(grid))
	for i := range l {
		l[i] = make([]bool, len(grid[i]))
	}
	return l
}

func (l landed) markLandAsLanded(i, j int, grid [][]byte) {
	if l[i][j] {
		return
	}
	l[i][j] = true
	l.markIfNotLanded(i-1, j, grid)
	l.markIfNotLanded(i+1, j, grid)
	l.markIfNotLanded(i, j-1, grid)
	l.markIfNotLanded(i, j+1, grid)
}

func (l landed) markIfNotLanded(i, j int, grid [][]byte) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return
	}
	if grid[i][j] == '1' && !l[i][j] {
		l.markLandAsLanded(i, j, grid)
	}
}

// @lc code=end
