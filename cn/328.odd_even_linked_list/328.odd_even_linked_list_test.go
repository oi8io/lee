package cn

import (
	"reflect"
	"testing"

	. "github.com/oi8io/lee/cn/common"
)

func Test_oddEvenList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"", args{BuildListNode([]int{1, 2, 3, 4, 5})}, BuildListNode([]int{1, 3, 5, 2, 4})},
		{"", args{BuildListNode([]int{1, 2, 3, 4})}, BuildListNode([]int{1, 3, 2, 4})},
		{"", args{BuildListNode([]int{1, 2, 3})}, BuildListNode([]int{1, 3, 2})},
		{"", args{BuildListNode([]int{1, 2})}, BuildListNode([]int{1, 2})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oddEvenList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("oddEvenList() = %v, want %v", got, tt.want)
			}
		})
	}
}
