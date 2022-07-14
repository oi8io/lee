package cn

import "testing"

func Test_longestValidParentheses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{"(()"}, 2},
		{"", args{"())"}, 2},
		{"", args{")()"}, 2},
		{"", args{")()())"}, 4},
		{"", args{")()()("}, 4},
		{"", args{""}, 0},
		{"", args{")))))))"}, 0},
		{"", args{")))))))(((((((("}, 0},
		{"", args{")))))))"}, 0},
		{"", args{"()"}, 2},
		{"", args{"()(()"}, 2},
		{"", args{"()((((((((())))))))"}, 16},
		{"", args{"()((((((((())))))))"}, 16},
		{"", args{"())()"}, 2},
		{"", args{"(()(((()"}, 2},
		{"", args{"(()(((()"}, 2},
		{"", args{")))()(()))())(())()())(()((())))())))))(())))(()()))(())())())))(()))()))((((()())))))()()))(()((())((())())()()()()((()((((())))(()))(()()()))))(()((((()))(((((()))())()))((("}, 52},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.args.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
		})
	}
}
