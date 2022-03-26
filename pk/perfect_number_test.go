package pk

import "testing"

func TestGetNumbers(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
	}{
		{"24", args{24}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetNumbers(tt.args.n)
		})
	}
}
