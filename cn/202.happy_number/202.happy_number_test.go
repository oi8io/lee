package cn

import "testing"

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{1}, true},
		{"", args{2}, true},
		{"", args{3}, true},
		{"", args{4}, true},
		{"", args{5}, true},
		{"", args{6}, true},
		{"", args{7}, true},
		{"", args{8}, true},
		{"", args{9}, true},
		{"", args{10}, true},
		{"", args{10000000}, true},
		{"", args{9999999999}, true},
		{"", args{19}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy(tt.args.n); got != tt.want {
				t.Errorf("isHappy() = %v, want %v", got, tt.want)
			}
		})
	}
}
