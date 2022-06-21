package cn

import "testing"

func Test_uniquePathsWithObstacles(t *testing.T) {
	type args struct {
		obstacleGrid [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		// 0000000010000
		// 0000000010000
		{"", args{[][]int{{1}, {0}}}, 0},
		{"", args{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsWithObstacles(tt.args.obstacleGrid); got != tt.want {
				t.Errorf("uniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
