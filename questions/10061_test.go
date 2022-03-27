package questions

import (
	"reflect"
	"testing"
)

func Test_rotateRight(t *testing.T) {
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
		{"", args{buildList([]int{1, 2, 3, 4, 5}), 2}, buildList([]int{4, 5, 1, 2, 3})},
		{"", args{buildList([]int{1, 2, 3, 4, 5, 6, 7, 8}), 3}, buildList([]int{6, 7, 8, 1, 2, 3, 4, 5})},
		{"", args{buildList([]int{1, 2, 3, 4, 5, 6, 7, 8}), 7}, buildList([]int{2, 3, 4, 5, 6, 7, 8, 1})},
		{"", args{buildList([]int{1, 2}), 1}, buildList([]int{2, 1})},
		{"", args{buildList([]int{2}), 1000}, buildList([]int{2})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRight(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
