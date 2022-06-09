package cn

import (
	"testing"
)

func Test_consecutiveNumbersSum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{9}, 3},
		{"", args{10}, 2},
		{"", args{15}, 4},
		{"", args{30}, 1},
		{"", args{85}, 4},
		{"", args{18}, 1},
		{"", args{199}, 1},
		{"", args{}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := consecutiveNumbersSum(tt.args.n); got != tt.want {
				t.Errorf("consecutiveNumbersSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
