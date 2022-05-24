package problems

import "testing"

func Test_repeatedNTimes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{1, 2, 3, 3}}, 3},
		{"", args{[]int{2, 6, 2, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedNTimes(tt.args.nums); got != tt.want {
				t.Errorf("repeatedNTimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
