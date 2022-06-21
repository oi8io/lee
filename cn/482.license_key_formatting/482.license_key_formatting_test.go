package cn

import "testing"

func Test_licenseKeyFormatting(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"", args{"5F3Z-2e-9-w", 4}, "5F3Z-2E9W"},
		{"", args{"zz5F3Z-2e-9-w", 4}, "ZZ-5F3Z-2E9W"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := licenseKeyFormatting(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("licenseKeyFormatting() = %v, want %v", got, tt.want)
			}
		})
	}
}
