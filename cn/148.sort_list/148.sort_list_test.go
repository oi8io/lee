package cn

import (
	"reflect"
	"testing"
)
import . "github.com/oi8io/lee/cn/common"

func Test_sortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"", args{nil}, nil},
		{"", args{BuildListNode([]int{1, 6, 3, 7, 4, 4, 1})}, BuildListNode([]int{1, 1, 3, 4, 4, 6, 7})},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", got, tt.want)
			}
		})
	}
}
