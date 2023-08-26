package binarysearchtest

import (
	"sort"
	"testing"
)

func Test_Binary_Search_Asc_Array(t *testing.T) {
	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 6
	cmp := func(i int) int {
		if a[i] > x {
			return -1
		}
		if a[i] == x {
			return 0
		}
		return 1
	}
	t.Error(sort.Find(len(a), cmp))
}

// lower bound
// a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
// x := 6

// i := sort.Search(len(a), func(i int) bool { return a[i] <= x })

func Test_UpperBound(t *testing.T) {
	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 6

	i := sort.Search(len(a), func(i int) bool { return a[i] > x })
	t.Error(i)
}
