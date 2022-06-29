package cn

import (
	. "github.com/oi8io/lee/cn/common"
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	var args1 = make([]*ListNode, 0)
	var x = [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}
	for i := 0; i < len(x); i++ {
		l := BuildListNode(x[i])
		args1 = append(args1, l)
	}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"", args{args1}, BuildListNode([]int{1, 1, 2, 3, 4, 4, 5, 6})},
		{"", args{nil}, nil},
		{"", args{make([]*ListNode, 0)}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
