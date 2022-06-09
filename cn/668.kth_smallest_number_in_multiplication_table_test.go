package cn

import "testing"

func Test_findKthNumber(t *testing.T) {
	type args struct {
		m int
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{5, 5, 9}, 0},
		{"", args{2, 3, 6}, 6},
		{"", args{3, 3, 5}, 3},
		{"", args{9, 9, 9}, 0},
		{"", args{30000, 30000, 19888}, 9600},
		{"", args{10000000000000, 1, 9}, 9},
		{"", args{10000000000000, 10000000000000, 9}, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthNumber(tt.args.m, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("findKthNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
