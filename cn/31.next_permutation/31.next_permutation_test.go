package cn

import (
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"", args{[]int{1, 2, 3}}},
		{"", args{[]int{3, 2, 1}}},
		{"", args{[]int{1, 1, 5}}},
		{"", args{[]int{1, 20, 26, 1, 15, 29, 4, 29, 10, 9, 21, 7, 27, 11, 21, 5, 9, 7, 27, 16, 17, 3, 6, 5, 16, 23, 29, 14, 28, 21, 2, 29, 3, 29, 0, 18, 28, 5, 10, 9, 6, 23, 8, 25, 26, 21, 1, 5, 29, 28, 14, 8, 1, 20, 13, 10}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nextPermutation(tt.args.nums)
			ns := tt.args.nums
			ret := []int{1, 20, 26, 1, 15, 29, 4, 29, 10, 9, 21, 7, 27, 11, 21, 5, 9, 7, 27, 16, 17, 3, 6, 5, 16, 23, 29, 14, 28, 21, 2, 29, 3, 29, 0, 18, 28, 5, 10, 9, 6, 23, 8, 25, 26, 21, 1, 5, 29, 28, 14, 8, 10, 1, 13, 20}
			for i := 0; i < len(ns); i++ {
				if ret[i] != ns[i] {
					t.Errorf("want %v got %v", ret[i:], ns[i:])
					break
				}
			}
		})
	}
}
