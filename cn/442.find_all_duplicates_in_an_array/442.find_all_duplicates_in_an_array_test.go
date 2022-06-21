package cn

import (
	"reflect"
	"testing"
)

func Test_findDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{[]int{4, 3, 2, 7, 8, 2, 3, 1}}, []int{3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicates(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDuplicates1(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{[]int{4, 3, 2, 7, 8, 2, 3, 1}}, []int{3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDuplicates1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
