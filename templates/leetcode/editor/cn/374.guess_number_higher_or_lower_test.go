package problems

import "testing"

func Test_guessNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{1000}, 9},
		{"", args{7}, 9},
		{"", args{10}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := guessNumber(tt.args.n); got != tt.want {
				t.Errorf("guessNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
