package cn

import "testing"
import . "github.com/oi8io/lee/cn/common"

func Test_kthSmallest(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{GetTreeByString("3,1,4,null,2"), 1}, 1},
		{"", args{GetTreeByString("3,1,4,null,2"), 2}, 2},
		{"", args{GetTreeByString("3,1,4,null,2"), 3}, 3},
		{"", args{GetTreeByString("3,1,4,null,2"), 4}, 4},
		{"", args{GetTreeByString("5,3,6,2,4,null,null,1"), 1}, 1},
		{"", args{GetTreeByString("5,3,6,2,4,null,null,1"), 2}, 2},
		{"", args{GetTreeByString("5,3,6,2,4,null,null,1"), 3}, 3},
		{"", args{GetTreeByString("5,3,6,2,4,null,null,1"), 4}, 4},
		{"", args{GetTreeByString("5,3,6,2,4,null,null,1"), 5}, 5},
		{"", args{GetTreeByString("5,3,6,2,4,null,null,1"), 6}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
