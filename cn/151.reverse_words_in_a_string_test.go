package cn

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"", args{"i1 i2 i3 i4          i5 i6 pneumonoultramicroscopicsilicovolcanoconiosis5 pneumonoultramicroscopicsilicovolcanoconiosis pneumonoultramicroscopicsilicovolcanoconiosis4 pneumonoultramicroscopicsilicovolcanoconiosis3 pneumonoultramicroscopicsilicovolcanoconiosis2 pneumonoultramicroscopicsilicovolcanoconiosis1"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
