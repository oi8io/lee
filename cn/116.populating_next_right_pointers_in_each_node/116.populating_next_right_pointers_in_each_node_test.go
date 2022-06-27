package cn

import (
	"reflect"
	"testing"
)

func Test_connect(t *testing.T) {
	type args struct {
		root *Node
	}

	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
		{"", args{&Node{Val: 1, Left: &Node{Val: 2}, Right: &Node{Val: 3}}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := connect(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connect() = %v, want %v", got, tt.want)
			}
		})
	}
}
