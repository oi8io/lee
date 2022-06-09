package cn

import (
	"reflect"
	"testing"
)

func Test_inorderSuccessor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
	}
	tree := &TreeNode{
		Val:  2,
		Left: nil,
		Right: &TreeNode{
			Val: 3,
		},
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		//Wrong Answer: input:[2,null,3] 2 Output:null Expected:3 stdout:
		{"", args{tree, tree}, tree.Right},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderSuccessor(tt.args.root, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderSuccessor() = %v, want %v", got, tt.want)
			}
		})
	}
}
