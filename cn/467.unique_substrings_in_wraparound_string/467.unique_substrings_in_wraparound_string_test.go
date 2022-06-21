package cn

import "testing"

func Test_findSubstringInWraproundString(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"ab", args{"a"}, 1},
		{"ab", args{"cac"}, 2},
		{"ab", args{"abaabaaba"}, 3},
		{"ab", args{"abcdcdefaaa"}, 3},
		{"ab", args{"abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxy"}, 3},
		{"ab", args{"zabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxy"}, 3},
		{"ab", args{"zazabcdefghttvuvxyzwwa"}, 3},
		{"ab", args{"zac"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstringInWraproundString(tt.args.p); got != tt.want {
				t.Errorf("findSubstringInWraproundString() = %v, want %v", got, tt.want)
			}
		})
	}
}
