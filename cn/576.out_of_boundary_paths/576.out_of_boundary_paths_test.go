package cn

import "testing"

func Test_findPaths(t *testing.T) {
	type args struct {
		m           int
		n           int
		maxMove     int
		startRow    int
		startColumn int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{2, 2, 2, 0, 0}, 6},
		{"", args{1, 3, 3, 0, 1}, 12},
		{"", args{8, 7, 16, 1, 5}, 102984580},
		{"", args{8, 50, 23, 5, 26}, 914783380},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPaths(tt.args.m, tt.args.n, tt.args.maxMove, tt.args.startRow, tt.args.startColumn); got != tt.want {
				t.Errorf("findPaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
