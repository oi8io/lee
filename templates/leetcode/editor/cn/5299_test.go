package problems

import "testing"

func Test_divisorSubstrings(t *testing.T) {
	type args struct {
		num int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{240, 2}, 2},
		{"", args{430043, 2}, 2},
		{"", args{680, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisorSubstrings(tt.args.num, tt.args.k); got != tt.want {
				t.Errorf("divisorSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
