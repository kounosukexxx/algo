package l1473

import (
	"testing"
)

func Test_minCost(t *testing.T) {
	type args struct {
		houses []int
		cost   [][]int
		m      int
		n      int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				houses: []int{0, 0, 0, 0, 0},
				cost: [][]int{
					{1, 10},
					{10, 1},
					{10, 1},
					{1, 10},
					{5, 1},
				},
				m:      5,
				n:      2,
				target: 3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.houses, tt.args.cost, tt.args.m, tt.args.n, tt.args.target); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minWithExcludeIndex(t *testing.T) {
	type args struct {
		a       [][]int
		exIndex int
		k       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				a: [][]int{
					{1, 10},
					{0, 1},
					{10, 1},
					{1, 10},
					{5, 1},
				},
				exIndex: 1,
				k:       0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWithExcludeIndex(tt.args.a, tt.args.exIndex, tt.args.k); got != tt.want {
				t.Errorf("minWithExcludeIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
