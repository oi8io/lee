package cn

import (
	cn "github.com/oi8io/lee/cn/common"
	"testing"
)

func Test_setZeroes(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"", args{[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cn.PrintMatrix(tt.args.matrix)
			setZeroes(tt.args.matrix)
			cn.PrintMatrix(tt.args.matrix)

		})
	}
}
