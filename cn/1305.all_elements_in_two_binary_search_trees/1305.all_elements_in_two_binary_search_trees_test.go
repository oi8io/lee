package cn

import (
	. "github.com/oi8io/lee/cn/common"
	"reflect"
	"testing"
)

func Test_getAllElements(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{GetTreeByString("1,2,4"), GetTreeByString("2,3,6")}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAllElements(tt.args.root1, tt.args.root2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAllElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
