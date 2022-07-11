package cn

import (
	"reflect"
	"testing"
)

func TestGetMatrix(t *testing.T) {
	type args struct {
		m int
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"", args{4, 5}, [][]int{}},
		{"", args{1, 5}, [][]int{}},
		{"", args{6, 3}, [][]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMatrix(tt.args.m, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintMatrix(t *testing.T) {
	PrintMatrix(GetMatrix(1, 5))
	PrintMatrix(GetMatrix(10, 1))
	PrintMatrix(GetMatrix(10, 3))
}
