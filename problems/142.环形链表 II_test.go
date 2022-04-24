package problems

import (
	"reflect"
	"testing"
)

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	list1 := BuildListNode([]int{1, 2})
	list1.Next.Next = list1
	list2 := BuildListNode([]int{3, 2, 0, -4})
	list2.Next.Next.Next.Next = list2.Next
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"", args{list1}, list1},
		{"", args{BuildListNode([]int{1})}, nil},
		{"", args{BuildListNode([]int{1, 2})}, nil},
		{"", args{list2}, list2.Next},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
