package cn

import (
	. "github.com/oi8io/lee/cn/common"
	"testing"
)

func Test_flatten(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"", args{GetTreeByString("1,2,5,3,4,null,6")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flatten(tt.args.root)
		})
	}
}
