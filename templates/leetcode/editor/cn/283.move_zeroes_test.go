package problems

import (
	"fmt"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"", args{[]int{0, 1}}},
		{"", args{[]int{1, 0}}},
		{"", args{[]int{1, 0, 0}}},
		{"", args{[]int{1, 0, 1}}},
		{"", args{[]int{0, 1, 0, 1}}},
		{"", args{[]int{0, 1, 0, 3, 12}}},
		{"", args{[]int{0, 1, 0, 3, 12}}},
		{"", args{[]int{0, 1, 0, 0, 0, 0, 3, 0, 0, 0, 0, 12}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
			fmt.Println(tt.args.nums)
		})
	}
}
