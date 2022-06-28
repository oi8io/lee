package cn

import (
	. "github.com/oi8io/lee/cn/common"
	"testing"
)

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{GetTreeByString("10,-9")}, 10},
		{"", args{GetTreeByString("-10,-9")}, -9},
		{"", args{GetTreeByString("[-10,9,20,null,null,15,7]")}, 42},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
