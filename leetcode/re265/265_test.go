package re265

import "testing"

func Test_getMins(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
		want2 int
	}{
		{
			name: "",
			args: args{
				a: []int{11, 2, 3},
			},
		},
		{
			name: "",
			args: args{
				a: []int{11, 2},
			},
		},
		{
			name: "",
			args: args{
				a: []int{11, 22},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := getMins(tt.args.a)
			if got != tt.want {
				t.Errorf("getMins() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getMins() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("getMins() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
