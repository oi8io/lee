package cn

import "testing"

func Test_totalNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{4}, 2},
		{"", args{8}, 92},
		{"", args{8}, 92},
		{"", args{8}, 92},
		{"", args{8}, 92},
		{"", args{8}, 92},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := totalNQueens(tt.args.n); got != tt.want {
				t.Errorf("totalNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_totalNQueens1(t *testing.T) {
	makeQueenMap(8)
}
