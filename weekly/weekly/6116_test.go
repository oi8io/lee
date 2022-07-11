package weekly

import "testing"
import . "github.com/oi8io/lee/cn/common"

func Test_evaluateTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{GetTreeByString("2,1,3,null,null,0,1")}, true},
		{"", args{GetTreeByString("0")}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evaluateTree(tt.args.root); got != tt.want {
				t.Errorf("evaluateTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
