package cn

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestBuildTreeWithString(t *testing.T) {
	type args struct {
		nodes string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{"", args{"1,2,3,null,5,6,null,8,9,11"}, nil},
		{"", args{"6,2,8,0,4,7,9,null,null,3,5"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildTreeWithString(tt.args.nodes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildTreeWithString() = %v, want %v", got, tt.want)
			}
		})
	}
}
func X() *TreeNode {
	str := "1,3,8,4,n,9,6,n,7,5,2"
	arr := strings.Split(str, ",")
	fmt.Println(arr)
	var x, i = 1, 0
	var aa = make([][]int, 0)
	var ar = make([]int, 0)
	var brk bool

	for {
		if brk {
			break
		}
		var a []int
		for j := x; j < x+x; j++ {
			if i == len(arr) {
				brk = true
				break
			}

			if checkParentIsNull(ar, j) {
				ar = append(ar, 0)
				continue
			}

			if arr[i] == "n" {
				i++
				a = append(a, 0)
				ar = append(ar, 0)
			} else {
				atoi, _ := strconv.Atoi(arr[i])

				a = append(a, atoi)
				ar = append(ar, atoi)
				i++
			}
			//fmt.Printf(" %d ", j)
		}
		fmt.Println()
		aa = append(aa, a)
		x = x * 2
	}
	root := &TreeNode{
		Val: ar[0],
	}
	for i := 1; i < len(ar)/2-1; i++ {
		if ar[i] == 0 {
			continue
		}
		if ar[2*i-1] != 0 {
			root.Left = &TreeNode{Val: 2*i - 1}
		}
		if ar[2*i] != 0 {
			root.Right = &TreeNode{Val: 2 * i}
		}
	}

	withArr := BuildTreeWithArr(ar, 1)
	return withArr
}

func TestVisitor(t *testing.T) {
	str := "1,3,8,4,n,9,6,n,7,5,2"
	arr := strings.Split(str, ",")
	fmt.Println(arr)
	var x, i = 1, 0
	var aa = make([][]int, 0)
	var ar = make([]int, 0)
	var brk bool

	for {
		if brk {
			break
		}
		var a []int
		for j := x; j < x+x; j++ {
			if i == len(arr) {
				brk = true
				break
			}

			if checkParentIsNull(ar, j) {
				ar = append(ar, 0)
				continue
			}

			if arr[i] == "n" {
				i++
				a = append(a, 0)
				ar = append(ar, 0)
			} else {
				atoi, _ := strconv.Atoi(arr[i])

				a = append(a, atoi)
				ar = append(ar, atoi)
				i++
			}
			//fmt.Printf(" %d ", j)
		}
		fmt.Println()
		aa = append(aa, a)
		x = x * 2
	}
	root := &TreeNode{
		Val: ar[0],
	}
	for i := 1; i < len(ar)/2-1; i++ {
		if ar[i] == 0 {
			continue
		}
		if ar[2*i-1] != 0 {
			root.Left = &TreeNode{Val: 2*i - 1}
		}
		if ar[2*i] != 0 {
			root.Right = &TreeNode{Val: 2 * i}
		}
	}
	fmt.Println(ar)
	withArr := BuildTreeWithArr(ar, 1)
	xxx := withArr.Val

	lx := levelOrder(withArr)
	fmt.Println(lx)
	fmt.Println(xxx)

	var k = 1
	fmt.Println()
	for {
		if len(ar) < k {
			break
		}
		for i := k; i < k+k && i <= len(ar); i++ {
			fmt.Printf(" %d ", ar[i-1])
		}
		fmt.Println()
		k = k * 2
	}
}

func TestBuildTreeWithString1(t *testing.T) {
	type args struct {
		nodes string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		// TODO: Add test cases.
		{"", args{"1,3,8,4,n,9,6,n,7,5,2"}, X()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildTreeWithString(tt.args.nodes)
			fmt.Println(levelOrder(got))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildTreeWithString() = %v, want %v", got, tt.want)
			}
		})
	}
}
