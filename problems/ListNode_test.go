package problems

import (
	"reflect"
	"testing"
)

func TestBuildListNode(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildListNode(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNode_String(t *testing.T) {

	tests := []struct {
		name   string
		fields *ListNode
		want   string
	}{
		// TODO: Add test cases.
		{"", BuildListNode([]int{1, 2, 3, 4, 5, 6}), " 1 , 2 "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := &ListNode{
				Val:  tt.fields.Val,
				Next: tt.fields.Next,
			}
			if got := head.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
