package basic

import "testing"

func TestFib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{3}, 2},
		{"", args{5}, 5},
		{"", args{8}, 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib(tt.args.n); got != tt.want {
				t.Errorf("Fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestFib1(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0", args{0}, 0},
		{"1", args{1}, 1},
		{"3", args{3}, 2},
		{"5", args{5}, 5},
		{"8", args{8}, 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib1(tt.args.n); got != tt.want {
				t.Errorf("Fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
