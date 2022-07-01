package cn

import (
	"reflect"
	"testing"

	. "github.com/oi8io/lee/cn/common"
)

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"", args{BuildListNode([]int{1})}, nil},
		{"", args{BuildListNode([]int{1, 2})}, nil},
		{"", args{BuildListNode([]int{1, 2, 3, 4, 5})}, nil},
		{"", args{BuildListNode([]int{1, 2, 3, 4, 5, 6})}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
