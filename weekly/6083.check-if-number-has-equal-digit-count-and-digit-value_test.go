package weekly

import "testing"

func Test_digitCount(t *testing.T) {
	type args struct {
		num string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{"1210"}, true},
		{"", args{"030"}, false},
		{"", args{"5210010007"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digitCount(tt.args.num); got != tt.want {
				t.Errorf("digitCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
