package cn

import "testing"

func Test_minMoves(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{5, 6, 8, 8, 5}}, 7},
		{"", args{[]int{5, 6, 8, 8, 5}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMoves(tt.args.nums); got != tt.want {
				t.Errorf("minMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
