package cn

import (
	. "github.com/oi8io/lee/cn/common"
	"testing"
)

func Test_reorderList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"", args{BuildListNode([]int{1, 2, 3, 4})}},
		{"", args{BuildListNode([]int{1, 2, 3, 4, 5})}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reorderList(tt.args.head)
		})
	}
}
