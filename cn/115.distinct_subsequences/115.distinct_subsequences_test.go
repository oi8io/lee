package cn

import "testing"

func Test_numDistinct(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{"b", "b"}, 1},
		{"", args{"a", "b"}, 0},
		{"", args{"aabb", "aab"}, 2},
		{"", args{"aabb", "abb"}, 2},
		{"", args{"aaaaaabb", "aaaaabbb"}, 0},
		{"", args{"aaaaaabb", "aaaaabb"}, 6},
		{"", args{"fff", "ffff"}, 0},
		{"", args{"adfassxfasdfa", "x"}, 1},
		{"", args{"aaaaaaaoxbbbbbbbbb", "xo"}, 0},
		{"", args{"aaa", "a"}, 3},
		{"", args{"aaa", "aa"}, 3},
		{"", args{"aaa", "aaa"}, 1},
		{"", args{"aaaa", "a"}, 4},
		{"", args{"aaaa", "aaa"}, 4},
		{"", args{"aaaa", "aa"}, 6},
		{"", args{"aaaaaa", "aaa"}, 20},
		{"", args{"", ""}, 0},
		{"", args{"babgbag", "bag"}, 5},
		{"", args{"rabbbit", "rabbit"}, 3},
		{"", args{"anacondastreetracecar", "contra"}, 6},
		{"", args{"condastreetraceca", "contra"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDistinct(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("numDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
