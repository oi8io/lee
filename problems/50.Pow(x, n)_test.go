package problems

import "testing"

func Test_myPow(t *testing.T) {
	type args struct {
		x float64
		n int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"2^10", args{2.0000, 10}, 1024.0000},
		{"2^-2", args{2.0000, -2}, 0.2500},
		{"2^7", args{2.0000, 7}, 128},
		{"2^27", args{2.0000, 27}, 100000000},
		{"1^10000000000", args{1.0000, 10000000000}, 1},
		{"1^10000000000", args{8.95371, -1}, 0.11169},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myPow(tt.args.x, tt.args.n); got != tt.want {
				t.Errorf("myPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
