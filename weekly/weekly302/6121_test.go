package weekly

import "testing"

func Test_sum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{36}, 9},
		{"", args{18}, 9},
		{"", args{43}, 7},
		{"", args{7}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sum(tt.args.n); got != tt.want {
				t.Errorf("sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{18, 43, 36, 13, 7}}, 54},
		{"", args{[]int{10, 12, 19, 14}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumSum(tt.args.nums); got != tt.want {
				t.Errorf("maximumSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
