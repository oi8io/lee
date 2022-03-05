package questions

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"abcabcbb", args{"abcabcbb"}, 3},
		{"bbbbb", args{"bbbbb"}, 1},
		{"pwwkew", args{"pwwkew"}, 3},
		{"abcdeazxvnmhjkl", args{"abcdeazxvnmhjkl"}, 14},
		{"tmmzuxt", args{"tmmzuxt"}, 5},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_lengthOfLongestSubstringBefore(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"abcabcbb", args{"abcabcbb"}, 3},
		{"bbbbb", args{"bbbbb"}, 1},
		{"pwwkew", args{"pwwkew"}, 3},
		{"abcdeazxvnmhjkl", args{"abcdeazxvnmhjkl"}, 14},
		{"tmmzuxt", args{"tmmzuxt"}, 5},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringBefore(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
