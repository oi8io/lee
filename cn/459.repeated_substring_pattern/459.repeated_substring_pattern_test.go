package cn

import "testing"

func Test_repeatedSubstringPattern(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{"aababcabcdabcdefgabce"}, false},
		{"", args{"abcdabcd"}, true},
		{"", args{"abcabc"}, true},
		{"", args{"abcabcabcabc"}, true},
		{"", args{"ababab"}, true},
		{"", args{"aaa"}, true},
		{"", args{"aa"}, true},
		{"", args{"aaaaaaaaaaaaa"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedSubstringPattern(tt.args.s); got != tt.want {
				t.Errorf("repeatedSubstringPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
