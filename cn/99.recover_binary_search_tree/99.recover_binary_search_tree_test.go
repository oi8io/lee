package cn

import "testing"
import . "github.com/oi8io/lee/cn/common"

func Test_recoverTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"", args{GetTreeByString("1,3,null,null,2")}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recoverTree(tt.args.root)
		})
	}
}
