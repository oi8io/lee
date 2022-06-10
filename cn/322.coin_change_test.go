package cn

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{1, 2, 5}, 11}, 3},
		{"", args{[]int{1, 2, 5}, 100}, 20},
		{"", args{[]int{1, 2, 5}, 9}, 3},
		{"", args{[]int{1, 2, 5}, 1}, 1},
		{"", args{[]int{186, 419, 83, 408}, 6249}, 20},
		{"", args{[]int{186, 419, 83, 408}, 5}, -1},
		{"", args{[]int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422}, 9864}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
