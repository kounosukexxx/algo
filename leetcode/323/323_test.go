package l323

import "testing"

func TestXxx(t *testing.T) {
	n := 5
	edges := [][]int{
		{0, 1},
		{1, 2},
		{3, 4},
	}
	t.Error(countComponents(n, edges))
}
