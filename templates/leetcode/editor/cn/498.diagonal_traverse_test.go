package problems

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_findDiagonalOrder(t *testing.T) {
	type args struct {
		mat [][]int
	}
	var x = make([][]int, 8)
	for i := 0; i < 8; i++ {
		x[i] = make([]int, 4)
		for j := 0; j < 4; j++ {
			x[i][j] = i*8 + j + 1
		}
	}
	fmt.Println(x)

	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, []int{1}},
		{"", args{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}}, []int{1}},
		{"", args{[][]int{{1, 2}, {3, 4}}}, []int{1}},
		{"", args{x}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDiagonalOrder(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDiagonalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
