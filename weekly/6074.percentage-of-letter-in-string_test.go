package weekly

import "testing"

func Test_percentageLetter(t *testing.T) {
	type args struct {
		s      string
		letter byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{"foobar", 'o'}, 33},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := percentageLetter(tt.args.s, tt.args.letter); got != tt.want {
				t.Errorf("percentageLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
