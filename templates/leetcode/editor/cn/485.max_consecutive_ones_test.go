package problems

import "testing"

func Test_findMaxConsecutiveOnes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{1, 1, 0, 1, 1, 1}}, 3},
		{"", args{[]int{1, 1, 0, 1}}, 2},
		{"", args{[]int{0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxConsecutiveOnes(tt.args.nums); got != tt.want {
				t.Errorf("findMaxConsecutiveOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
