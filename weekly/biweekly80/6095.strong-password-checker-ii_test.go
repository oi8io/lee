package biweekly80

import "testing"

func Test_strongPasswordCheckerII(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{"!IloveLe3tcode"}, true},
		{"", args{"Me+You--IsMyDream"}, false},
		{"", args{"1aB!"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strongPasswordCheckerII(tt.args.password); got != tt.want {
				t.Errorf("strongPasswordCheckerII() = %v, want %v", got, tt.want)
			}
		})
	}
}
