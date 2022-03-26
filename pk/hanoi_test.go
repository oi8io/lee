package pk

import (
	"reflect"
	"testing"
)

func TestHanoi(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"4", args{4}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hanoi(tt.args.n); got != tt.want {
				t.Errorf("Hanoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTowerOfHanoi(t *testing.T) {
	type args struct {
		n int
		a string
		b string
		c string
	}
	tests := []struct {
		name string
		args args
	}{
		{"4", args{4, "A", "B", "C"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TowerOfHanoi(tt.args.n, tt.args.a, tt.args.b, tt.args.c)
		})
	}
}

func TestTowerOfHanoi2(t *testing.T) {
	type args struct {
		a []int
		b []int
		c []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"4", args{[]int{4, 3, 2, 1}, []int{}, []int{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TowerOfHanoi2(tt.args.a, tt.args.b, tt.args.c)
		})
	}
}

func TestHanota(t *testing.T) {
	type args struct {
		A []int
		B []int
		C []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"4", args{[]int{2, 1, 0}, []int{}, []int{}}, []int{2, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hanota(tt.args.A, tt.args.B, tt.args.C); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Hanota() = %v, want %v", got, tt.want)
			}
		})
	}
}
