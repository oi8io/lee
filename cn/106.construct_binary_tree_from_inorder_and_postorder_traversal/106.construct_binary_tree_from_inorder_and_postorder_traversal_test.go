package cn

import (
	"reflect"
	"testing"
	
	. "github.com/oi8io/lee/cn/common"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		inorder   []int
		postorder []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{"", args{[]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}}, nil},
		{"", args{[]int{1, 2}, []int{2, 1}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildTree(tt.args.inorder, tt.args.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
