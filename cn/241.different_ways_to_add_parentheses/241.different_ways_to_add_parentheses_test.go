package cn

import (
	"reflect"
	"sort"
	"testing"
)

func Test_diffWaysToCompute(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{"2*3-4"}, []int{-34, -10, -14, -10, 10}},
		{"", args{"2*3*4*5"}, []int{-34, -10, -14, -10, 10}},
		{"", args{"2*1*15"}, []int{1, 2, 3}},
		{"", args{"2*3-4*5*3"}, []int{-114, -114, -30, -102, -30, -54, -54, -30, 30, -102, -30, -42, -30, 30}},
		{"", args{"2*3-4*5*3-4*5"}, []int{1, 2, 3}},
		{"", args{"2*3-4*5*3-4*5*3-4*5"}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := diffWaysToCompute(tt.args.expression)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diffWaysToCompute() = %v, want %v", got, tt.want)
			}
		})
	}
}
