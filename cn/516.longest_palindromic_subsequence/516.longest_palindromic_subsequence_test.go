package cn

import "testing"

func Test_longestPalindromeSubseq(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{"ab"}, 1},
		{"", args{"bb"}, 2},
		{"", args{"bab"}, 3},
		{"", args{"bbb"}, 3},
		{"", args{"bbbb"}, 4},
		{"", args{"bbbbb"}, 5},
		{"", args{"abcdefgh"}, 1},
		{"", args{"abcdefghh"}, 2},
		{"", args{"bbbbbbbbb"}, 9},
		{"", args{"bbbab"}, 4},
		{"", args{"babbb"}, 4},
		{"", args{"babbb"}, 4},
		{"", args{"cdbbcccabbda"}, 9},
		{"", args{"aaa"}, 3},
		{"", args{"euazbipzncptldueeuechubrcourfpftcebikrxhybkymimgvldiwqvkszfycvqyvtiwfckexmowcxztkfyzqovbtmzpxojfofbvwnncajvrvdbvjhcrameamcfmcoxryjukhpljwszknhiypvyskmsujkuggpztltpgoczafmfelahqwjbhxtjmebnymdyxoeodqmvkxittxjnlltmoobsgzdfhismogqfpfhvqnxeuosjqqalvwhsidgiavcatjjgeztrjuoixxxoznklcxolgpuktirmduxdywwlbikaqkqajzbsjvdgjcnbtfksqhquiwnwflkldgdrqrnwmshdpykicozfowmumzeuznolmgjlltypyufpzjpuvucmesnnrwppheizkapovoloneaxpfinaontwtdqsdvzmqlgkdxlbeguackbdkftzbnynmcejtwudocemcfnuzbttcoew"}, 159},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindromeSubseq(tt.args.s); got != tt.want {
				t.Errorf("longestPalindromeSubseq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonSubsequence516(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{"oxcpqrsvwf", "shmtulqrypy"}, 2},
		{"", args{"abcde", "ace"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequence516(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("longestCommonSubsequence516() = %v, want %v", got, tt.want)
			}
		})
	}
}
