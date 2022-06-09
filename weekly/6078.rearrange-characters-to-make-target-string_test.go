package weekly

import "testing"

func Test_rearrangeCharacters(t *testing.T) {
	type args struct {
		s      string
		target string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{"ilovecodingonleetcode", "code"}, 2},
		{"", args{"abcba", "abc"}, 1},
		{"", args{"abbaccaddaeea", "aaaaa"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rearrangeCharacters(tt.args.s, tt.args.target); got != tt.want {
				t.Errorf("rearrangeCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
