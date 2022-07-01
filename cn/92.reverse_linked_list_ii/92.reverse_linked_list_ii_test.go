package cn

import (
	"reflect"
	"testing"

	. "github.com/oi8io/lee/cn/common"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"", args{BuildListNode([]int{1, 9, 3, 6, 2, 4, 5}), 2, 5}, nil},
		{"", args{BuildListNode([]int{1, 2, 3, 4, 5, 6, 7}), 2, 5}, nil},
		{"", args{BuildListNode([]int{1, 2, 3, 4, 5, 6, 7}), 1, 7}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
