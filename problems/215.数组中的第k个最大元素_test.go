package problems

import "testing"

func Test_findKthLargest(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{2, 3}, 1}, 3},
		{"", args{[]int{2, 3}, 2}, 2},
		{"", args{[]int{1, 2, 3}, 1}, 3},
		{"", args{[]int{3, 2, 1, 5, 6, 4}, 1}, 6},
		{"", args{[]int{3, 2, 1, 5, 6, 4}, 2}, 5},
		{"", args{[]int{3, 2, 1, 5, 6, 4}, 3}, 4},
		{"", args{[]int{3, 2, 1, 5, 6, 4}, 4}, 3},
		{"", args{[]int{3, 2, 1, 5, 6, 4}, 5}, 2},
		{"", args{[]int{3, 2, 1, 5, 6, 4}, 6}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
