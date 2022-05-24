package problems

import "testing"

func Test_isAlienSorted(t *testing.T) {
	type args struct {
		words []string
		order string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		//输入：words = ["hello","leetcode"], order = "hlabcdefgijkmnopqrstuvwxyz"
		{"", args{[]string{"hello", "leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"}, true},
		//输入：words = ["apple","app"], order = "abcdefghijklmnopqrstuvwxyz"
		{"", args{[]string{"apple", "app"}, "abcdefghijklmnopqrstuvwxyz"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAlienSorted(tt.args.words, tt.args.order); got != tt.want {
				t.Errorf("isAlienSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
