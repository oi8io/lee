package cn

import (
	"reflect"
	"testing"
)
import . "github.com/oi8io/lee/cn/common"

func Test_sortedListToBST(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{"", args{BuildListNode([]int{-10, -3, 0, 5, 9})}, GetTreeByString("0,-3,9,-10,null,5")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedListToBST(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedListToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
