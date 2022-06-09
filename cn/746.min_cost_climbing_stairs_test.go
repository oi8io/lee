package cn

import "testing"

func Test_minCostClimbingStairs(t *testing.T) {
	type args struct {
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}}, 6},
		{"", args{[]int{1, 100, 1, 1, 1, 100, 1, 100, 1, 1}}, 6},
		{"", args{[]int{1, 10, 100, 10000, 1, 100, 1, 100, 1, 1}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs(tt.args.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
