package cn

import (
	"reflect"
	"testing"
)

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"", args{[]int{1, 2, 3}}, [][]int{{1, 2, 3}}},
		{"", args{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}}, [][]int{{1, 2, 3}}},
		{"", args{[]int{9, 0, 3, 5, 7}}, [][]int{{9, 0, 3, 5, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
