package rere560

import "testing"

func Test_subarraySum(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 1, 1},
				k:    2,
			},
		},
		{
			name: "1",
			args: args{
				nums: []int{1, 2, 3},
				k:    3,
			},
		},
		{
			name: "1",
			args: args{
				nums: []int{1, 1, -2, 2},
				k:    2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraySum(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("subarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
