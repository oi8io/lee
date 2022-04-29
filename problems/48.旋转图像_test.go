package problems

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{"3x3", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}},
		{"4x4", args{[][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.matrix)
		})
	}
}

func Test_rotate1(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{"3x3", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate1(tt.args.matrix)
		})
	}
}

func Test_rotate11(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"3x3", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		{"4x4", args{[][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}}, [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotate1(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotate1() = %v, want %v", got, tt.want)
			}
		})
	}
}
