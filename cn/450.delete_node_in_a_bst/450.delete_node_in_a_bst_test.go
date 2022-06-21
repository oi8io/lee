package cn

import (
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	type args struct {
		root *TreeNode
		key  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{"", args{BuildTreeWithString("5,3,6,2,4,null,7"), 3}, BuildTreeWithString1("5,4,6,2,null,null,7")},
		{"", args{BuildTreeWithString("5,3,6,2,4,null,7"), 5}, BuildTreeWithString1("6,3,7,2,4")},
		{"", args{BuildTreeWithString("5,3,6,2,4,null,7"), 0}, BuildTreeWithString1("5,3,6,2,4,null,7")},
		{"", args{BuildTreeWithString(""), 0}, BuildTreeWithString1("")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(tt.args.root, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
