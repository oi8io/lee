package cn

import "testing"

func Test_numPrimeArrangements(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{100}, 0},
		{"", args{99}, 0},
		{"", args{98}, 892915734},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numPrimeArrangements(tt.args.n); got != tt.want {
				t.Errorf("numPrimeArrangements() = %v, want %v", got, tt.want)
			}
		})
	}
}
