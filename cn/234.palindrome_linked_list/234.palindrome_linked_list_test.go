package cn

import "testing"
import . "github.com/oi8io/lee/cn/common"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{BuildListNode([]int{1})}, true},
		{"", args{BuildListNode([]int{1, 1})}, true},
		{"", args{BuildListNode([]int{1, 2})}, false},
		{"", args{BuildListNode([]int{1, 2, 2, 1})}, true},
		{"", args{BuildListNode([]int{1, 2, 3, 2, 1})}, true},
		{"", args{BuildListNode([]int{1, 2, 5, 4, 2, 1})}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
