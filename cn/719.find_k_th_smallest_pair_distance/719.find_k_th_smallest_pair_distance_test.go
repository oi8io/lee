package cn

import "testing"

func Test_smallestDistancePair(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	var nums1 []int
	for i := 0; i < 10; i++ {
		nums1 = append(nums1, i)
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{nums1, 8888}, 1},
		{"", args{[]int{1, 1, 1}, 1}, 0},
		{"", args{[]int{1, 6, 1}, 1}, 0},
		{"", args{[]int{1, 6, 1}, 3}, 5},
		{"", args{[]int{1, 6, 1}, 3}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestDistancePair(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("smallestDistancePair() = %v, want %v", got, tt.want)
			}
		})
	}
}
