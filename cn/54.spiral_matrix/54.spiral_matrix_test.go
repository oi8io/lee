package cn

import (
	. "github.com/oi8io/lee/cn/common"
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{GetMatrix(1, 5)}, []int{1, 2, 3, 4, 5}},
		{"", args{GetMatrix(5, 2)}, []int{1, 2, 4, 6, 8, 10, 9, 7, 5, 3}},
		{"", args{GetMatrix(2, 5)}, []int{1, 2, 3, 4, 5, 10, 9, 8, 7, 6}},
		{"", args{GetMatrix(10, 1)}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"", args{GetMatrix(10, 3)}, []int{1, 2, 3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 29, 28, 25, 22, 19, 16, 13, 10, 7, 4, 5, 8, 11, 14, 17, 20, 23, 26}},
		{"", args{GetMatrix(3, 10)}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 20, 30, 29, 28, 27, 26, 25, 24, 23, 22, 21, 11, 12, 13, 14, 15, 16, 17, 18, 19}},
		{"", args{GetMatrix(3, 9)}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 18, 27, 26, 25, 24, 23, 22, 21, 20, 19, 10, 11, 12, 13, 14, 15, 16, 17}},
		{"", args{[][]int{{1, 2, 3, 4, 5}}}, []int{1, 2, 3, 4, 5}},
		{"", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{"", args{[][]int{{1, 2}, {4, 3}}}, []int{1, 2, 3, 4}},
		{"", args{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := spiralOrder(tt.args.matrix)
			PrintMatrix(tt.args.matrix)
			PrintArray(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
