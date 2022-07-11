package weekly

import "testing"

func Test_idealArrays(t *testing.T) {
	type args struct {
		n        int
		maxValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{2, 5}, 10},
		{"", args{5, 3}, 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := idealArrays(tt.args.n, tt.args.maxValue); got != tt.want {
				t.Errorf("idealArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
