package questions

import (
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"", args{buildList([]int{1, 2, 3, 4, 5}), 2}, buildList([]int{2, 1, 4, 3, 5})},
		{"", args{buildList([]int{1, 2}), 2}, buildList([]int{2, 1})},
		{"", args{buildList([]int{1, 2, 3, 4, 5}), 3}, buildList([]int{3, 2, 1, 4, 5})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListNode_String(t *testing.T) {
	type fields struct {
		Next *ListNode
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
		{"", fields{buildList([]int{1, 2, 3, 4, 5})}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ListNode{
				Next: tt.fields.Next,
			}
			if got := l.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildList(t *testing.T) {
	type args struct {
		in []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"", args{[]int{1, 2, 3, 4, 5}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildList(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildList() = %v, want %v", got, tt.want)
			}
		})
	}
}
