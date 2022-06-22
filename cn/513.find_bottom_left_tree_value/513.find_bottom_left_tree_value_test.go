package cn

import (
	cn "github.com/oi8io/lee/cn/tree"
	"testing"
)

func Test_findBottomLeftValue(t *testing.T) {
	type args struct {
		root *cn.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{cn.BuildTreeWithString("2,1,3")}, 1},
		{"", args{cn.BuildTreeWithString("1,2,3,4,null,5,6,null,null,7")}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBottomLeftValue(tt.args.root); got != tt.want {
				t.Errorf("findBottomLeftValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
