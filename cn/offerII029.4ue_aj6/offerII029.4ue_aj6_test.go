package cn

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func Test_insert(t *testing.T) {
	type args struct {
		aNode *Node
		x     int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		// TODO: Add test cases.
		{"", args{MakeCycleLink([]int{1, 2}), 3}, MakeCycleLink([]int{1, 2, 3})},
		{"", args{MakeCycleLink([]int{1, 3, 5}), 5}, MakeCycleLink([]int{1, 3, 5, 5})},
		{"", args{MakeCycleLink([]int{3, 4}), 2}, MakeCycleLink([]int{4, 2, 3})},
		{"", args{MakeCycleLink([]int{0, 0, 0}), 2}, MakeCycleLink([]int{0, 0, 0, 2})},
		{"", args{MakeCycleLink([]int{0, 0, 0}), 0}, MakeCycleLink([]int{0, 0, 0, 0})},
		{"", args{MakeCycleLink([]int{1}), 1}, MakeCycleLink([]int{1, 1})},
		{"", args{nil, 2}, MakeCycleLink([]int{2})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := insert(tt.args.aNode, tt.args.x)
			PrintNodes(got)
			PrintNodes(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func MakeCycleLink(in []int) *Node {
	if len(in) == 0 {
		return nil
	}
	var h = &Node{Val: in[0]}
	nh := h
	for i := 1; i < len(in); i++ {
		nh.Next = &Node{Val: in[i]}
		nh = nh.Next
	}
	nh.Next = h
	return nh
}
func PrintNodes(x *Node) {
	if x == nil {
		fmt.Println("[]")
		return
	}
	fmt.Print("[")
	h := x
	var arr []string
	for {
		arr = append(arr, strconv.Itoa(h.Val))
		h = h.Next
		if h == x {
			break
		}
	}
	fmt.Print(strings.Join(arr, "->"))
	fmt.Println("]")
}
