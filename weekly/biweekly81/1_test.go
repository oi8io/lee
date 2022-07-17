package weekly

import "testing"

func Test_countAsterisks(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{"l|*e*et|c**o|*de|"}, 2},
		{"", args{"iamprogrammer"}, 0},
		{"", args{"yo|uar|e**|b|e***au|tifu|l"}, 5},
		{"", args{"*||    *    ||    ||     |*|    *      |***|     |*|    |***|"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAsterisks(tt.args.s); got != tt.want {
				t.Errorf("countAsterisks() = %v, want %v", got, tt.want)
			}
		})
	}
}
