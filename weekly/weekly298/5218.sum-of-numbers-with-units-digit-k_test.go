package weekly

import "testing"

func Test_minimumNumbers(t *testing.T) {
	type args struct {
		num int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{58, 9}, 2},
		{"", args{37, 2}, -1},
		{"", args{0, 7}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumNumbers(tt.args.num, tt.args.k); got != tt.want {
				t.Errorf("minimumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
