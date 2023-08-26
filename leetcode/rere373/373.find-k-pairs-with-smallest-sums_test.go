package rere373

import (
	"reflect"
	"testing"
)

func Test_kSmallestPairs(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				nums1: []int{1, 7, 11},
				nums2: []int{2, 4, 6},
				k:     9,
			},
		},
		{
			name: "2",
			args: args{
				nums1: []int{1, 1, 2},
				nums2: []int{1, 2, 3},
				k:     9,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kSmallestPairs(tt.args.nums1, tt.args.nums2, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kSmallestPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
