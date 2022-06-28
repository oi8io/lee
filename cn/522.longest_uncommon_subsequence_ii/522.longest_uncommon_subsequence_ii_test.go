package cn

import "testing"

func Test_findLUSlength(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]string{"aaaaaaaaaaaaaaa", "a", "aaaa", "aa", "aaaaaaa"}}, 15},
		{"", args{[]string{"aaa", "acb"}}, 3},
		{"", args{[]string{"aabbcc", "aabbcc", "cb"}}, 2},
		{"", args{[]string{"ab", "abc"}}, 3},
		{"", args{[]string{"ac", "abc", "abc"}}, -1},
		{"", args{[]string{"aba", "cdc", "eae"}}, 3},
		{"", args{[]string{"abcdefgh", "hgfedcba"}}, 8},
		{"", args{[]string{"a", "b", "c", "d", "e", "f", "a", "b", "c", "d", "e", "f"}}, -1},
		{"", args{[]string{"aabbcc", "aabbcc", "c", "e", "aabbcd"}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLUSlength(tt.args.strs); got != tt.want {
				t.Errorf("findLUSlength() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSubSeq(t *testing.T) {
	type args struct {
		s   string
		sub string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{"aabbcc", "cb"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubSeq(tt.args.s, tt.args.sub); got != tt.want {
				t.Errorf("isSubSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}
