package problems

import (
	"reflect"
	"testing"
)

func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"", args{BuildListNode([]int{2}), 2}, BuildListNode([]int{})},
		{"", args{BuildListNode([]int{1, 2, 6, 3, 4, 5, 6}), 6}, BuildListNode([]int{1, 2, 3, 4, 5})},
		{"", args{BuildListNode([]int{6, 6, 1, 2, 6, 3, 4, 5, 6}), 6}, BuildListNode([]int{1, 2, 3, 4, 5})},
		{"", args{BuildListNode([]int{1, 2, 2, 1}), 2}, BuildListNode([]int{1, 1})},
		{"", args{BuildListNode([]int{5, 5, 5, 5, 5, 5, 5}), 5}, BuildListNode([]int{})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElements(tt.args.head, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
