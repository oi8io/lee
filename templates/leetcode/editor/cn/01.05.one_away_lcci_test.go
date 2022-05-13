package problems

import "testing"

func Test_oneEditAway(t *testing.T) {
	type args struct {
		first  string
		second string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"ab", args{"a", "b"}, true},
		{"ABC", args{"AB", "BC"}, false},
		{"Ple", args{"Ple", "Pale"}, true},
		{"Pzle", args{"Pzle", "Pale"}, true},
		{"horse", args{"horse", "ros"}, false},
		{"horse", args{"teacher", "teach"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := oneEditAway(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("oneEditAway() = %v, want %v", got, tt.want)
			}
		})
	}
}
