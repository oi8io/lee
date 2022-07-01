package cn

import (
	"reflect"
	"testing"

	. "github.com/oi8io/lee/cn/common"
)

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	//[4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2,
	a := BuildListNode([]int{4, 1})
	b := BuildListNode([]int{5, 6, 1})
	c := BuildListNode([]int{8, 4, 5})
	a.GetLast().Next = c
	b.GetLast().Next = c
	Dump(a)
	Dump(b)
	Dump(c)
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"", args{a, b}, c},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
