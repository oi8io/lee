package problems

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		//* 输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
		//* 输出：[[1,6],[8,10],[15,18]]
		{"", args{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{"", args{[][]int{{1, 4}, {4, 5}}}, [][]int{{1, 5}}},
		{"", args{[][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}}, [][]int{{1, 10}}},
		{"", args{[][]int{{2, 3}, {5, 5}, {2, 2}, {3, 4}, {3, 4}}}, [][]int{{2, 4}, {5, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
