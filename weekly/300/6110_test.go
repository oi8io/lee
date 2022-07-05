package weekly

import (
	. "github.com/oi8io/lee/cn/common"
	"reflect"
	"testing"
)

func Test_spiralMatrix(t *testing.T) {
	type args struct {
		m    int
		n    int
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"", args{3, 5, BuildListNode([]int{3, 0, 2, 6, 8, 1, 7, 9, 4, 2, 5, 5, 0})}, [][]int{{3, 0, 2, 6, 8}, {5, 0, -1, -1, 1}, {5, 2, 4, 9, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralMatrix(tt.args.m, tt.args.n, tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
