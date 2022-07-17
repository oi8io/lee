package weekly

import "testing"

func Test_longestSubsequence(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{"1001010", 5}, 5},
		{"", args{"00101001", 1}, 6},
		{"", args{"00101001", 1}, 6},
		{"", args{"11111111", 1}, 1},
		{"", args{"11111111", 5}, 2},
		{"", args{"11111111", 5}, 2},
		{"", args{"11111111", 0}, 0},
		{"", args{"0", 583196182}, 1},
		{"", args{"001010101011010100010101101010010", 93951055}, 31},
		{"", args{"001010101011010100010101101010010", 5}, 19},
		{"", args{"111100010000011101001110001111000000001011101111111110111000011111011000010101110100110110001111001001011001010011010000011111101001101000000101101001110110000111101011000101", 11713332}, 174},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubsequence(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("longestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
