package cn

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{"sssfasfsaf", ""}, 0},
		{"", args{"", ""}, 0},
		{"", args{"", "1"}, -1},
		{"", args{"aaaa", "aaa"}, 0},
		{"", args{"aaaa", "baa"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
