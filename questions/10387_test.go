package questions

import "testing"

func TestFirstUniqChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"",args{"abcaadbscde"},7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstUniqChar(tt.args.s); got != tt.want {
				t.Errorf("FirstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
