package cn

import "testing"

func Test_isMatch(t *testing.T) {
	type args struct {
		s string
		p string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{"a", "a"}, true},
		{"", args{"aaaa", "a*"}, true},
		{"", args{"aaaaab", "a*b"}, true},
		{"", args{"aaaaab", "a*"}, false},
		{"", args{"ab", ".*c"}, false},
		{"", args{"abc", ".*c"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
