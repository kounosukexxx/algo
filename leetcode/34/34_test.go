package l34

import (
	"reflect"
	"testing"
)

func Test_reverseSearch(t *testing.T) {
	nums := []int{1, 2, 3, 3, 3, 5, 6}
	nums2 := []int{1}
	type args struct {
		n int
		f func(int) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "4",
			args: args{
				n: len(nums),
				f: func(i int) bool {
					return nums[i] <= 4
				},
			},
			want: 4,
		},
		{
			name: "1",
			args: args{
				n: len(nums2),
				f: func(i int) bool {
					return nums2[i] <= 1
				},
			},
			want: 0,
		},
		{
			name: "not found",
			args: args{
				n: len(nums),
				f: func(i int) bool {
					return nums[i] <= 0
				},
			},
			want: -1,
		},
		{
			name: "got 0",
			args: args{
				n: len(nums),
				f: func(i int) bool {
					return nums[i] <= 1
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseSearch(tt.args.n, tt.args.f); got != tt.want {
				t.Errorf("reverseSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchRange(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: []int{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
