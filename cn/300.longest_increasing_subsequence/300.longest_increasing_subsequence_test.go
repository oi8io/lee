package cn

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{10, 9, 2, 5, 3, 7, 101, 18}}, 4},
		{"", args{[]int{10, 9, 2, 5, 3, 7, 101, 18, 2}}, 4},
		{"", args{[]int{7, 7, 7, 7, 7, 7, 7}}, 1},
		{"", args{[]int{0, 1, 0, 3, 2, 3}}, 4},
		{"", args{[]int{4, 10, 4, 3, 8, 9}}, 3},
		{"", args{[]int{4}}, 1},
		{"", args{[]int{10, 4}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
