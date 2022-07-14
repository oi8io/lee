package cn

import (
	"reflect"
	"testing"
)

func Test_asteroidCollision(t *testing.T) {
	type args struct {
		asteroids []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{[]int{5, 10, -5}}, []int{5, 10}},
		{"", args{[]int{10, 2, -5}}, []int{10}},
		{"", args{[]int{-2, -1, 1, 2}}, []int{-2, -1, 1, 2}},
		{"", args{[]int{1, -2, -2, -2}}, []int{-2, -2, -2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := asteroidCollision(tt.args.asteroids); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("asteroidCollision() = %v, want %v", got, tt.want)
			}
		})
	}
}
