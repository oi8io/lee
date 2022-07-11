package cn

import (
	cn "github.com/oi8io/lee/cn/common"
	"testing"
)

func Test_cherryPickup(t *testing.T) {
	type args struct {
		grid [][]int
	}
	var g = make([][]int, 50)
	for i := 0; i < 50; i++ {
		g[i] = make([]int, 50)
		for j := 0; j < 50; j++ {
			g[i][j] = 1
		}
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{cn.GetMatrix(3, 3)}, 42},
		{"", args{g}, 196},
		{"", args{[][]int{{0, 1, -1}, {1, 0, -1}, {1, 1, 1}}}, 5},
		{"", args{[][]int{{0}}}, 0},
		{"", args{[][]int{{1}}}, 1},
		{"", args{[][]int{{-1}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cherryPickup(tt.args.grid); got != tt.want {
				t.Errorf("cherryPickup() = %v, want %v", got, tt.want)
			}
		})
	}
}
