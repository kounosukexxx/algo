package l45

import (
	"testing"
)

func Test_max(t *testing.T) {
	type args struct {
		nums    []int
		from    int
		jumpCap int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				nums:    []int{1, 2, 4, 3, 5},
				from:    1,
				jumpCap: 2,
			},
			want: 4,
		},
		{
			name: "1",
			args: args{
				nums:    []int{1, 2, 4, 3, 5},
				from:    1,
				jumpCap: 76,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toJump(tt.args.nums, tt.args.from, tt.args.jumpCap); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}
