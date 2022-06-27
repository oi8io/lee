package cn

import "testing"

func Test_isValidSerialization(t *testing.T) {
	type args struct {
		preorder string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{"9,3,4,#,#,1,#,#,2,#,6,#,#"}, true},
		{"", args{"1,#"}, false},
		{"", args{"1,#"}, false},
		{"", args{"9,#,#,1"}, false},
		{"", args{"#"}, true},
		{"", args{""}, true},
		{"", args{"1,2,3,#,#,#,#"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSerialization(tt.args.preorder); got != tt.want {
				t.Errorf("isValidSerialization() = %v, want %v", got, tt.want)
			}
		})
	}
}
