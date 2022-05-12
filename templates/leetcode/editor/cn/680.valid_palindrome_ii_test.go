package problems

import "testing"

func Test_validPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{"cbbc"}, true},
		{"", args{"abac"}, true},
		{"", args{"caba"}, true},
		{"", args{"abc"}, false},
		{"", args{"adadadadcccccabaaaaaaaaaaaaaaabcccccdadadada"}, true},
		{"", args{"abcdcbxa"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validPalindrome(tt.args.s); got != tt.want {
				t.Errorf("validPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
