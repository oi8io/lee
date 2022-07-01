package cn

import (
	"reflect"
	"testing"

	. "github.com/oi8io/lee/cn/common"
)

func Test_insertionSortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"", args{BuildListNode([]int{4, 2, 1, 3})}, BuildListNode([]int{1, 2, 3, 4})},
		{"", args{BuildListNode([]int{4, 2, 1, 3, 5})}, BuildListNode([]int{1, 2, 3, 4, 5})},
		{"", args{BuildListNode([]int{1, 2})}, BuildListNode([]int{1, 2})},
		{"", args{BuildListNode([]int{1})}, BuildListNode([]int{1})},
		{"", args{BuildListNode([]int{2, 2})}, BuildListNode([]int{2, 2})},
		{"", args{BuildListNode([]int{2, 1})}, BuildListNode([]int{1, 2})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertionSortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertionSortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
