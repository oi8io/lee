package cn

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{root: BuildTreeWithString("5,4,6,null,null,3,7")}, false},
		{"", args{root: BuildTreeWithString("2,1,3")}, true},
		{"", args{root: BuildTreeWithString("3,1,5,0,2,4,6")}, true},
		{"", args{root: BuildTreeWithString("2147483647,-2147483648")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
