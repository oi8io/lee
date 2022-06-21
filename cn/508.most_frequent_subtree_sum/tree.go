package cn

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Node struct {
	Val      int
	Children []*Node
}

func BuildTree(in []int) *TreeNode {
	root := &TreeNode{Val: in[0]}
	for i := 1; i < len(in); i++ {
		root.Left = &TreeNode{}
	}
	return nil
}

func (t *TreeNode) LeftCheck() {
	if t == nil {
		return
	}
	if t.Left != nil {
		t.Left.LeftCheck()
	}
	fmt.Println(t.Val)
	if t.Right != nil {
		t.Right.LeftCheck()
	}
}

func (t *TreeNode) Add(num int) {
	if t == nil {
		return
	}
	if t.Left == nil && t.Val != -1 {
		t.Left = &TreeNode{Val: num}
		return
	}
	//左边为实，右边为空 先填充右边
	if t.Left != nil {
		if t.Right == nil {
			t.Right = &TreeNode{Val: num}
		} else {
			t.Left.Add(num)
		}
	} else {
		if t.Left.Val != -1 {

		}
	}

}

func BuildTreeWithString1(nodes string) *TreeNode {
	arr := strings.Split(nodes, ",")
	var ar = make([]int, 0)
	var brk bool
	var x, i = 1, 0
	for {
		if brk {
			break
		}
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
				ar = append(ar, 0)
			} else {
				atoi, _ := strconv.Atoi(arr[i])
				ar = append(ar, atoi)
				i++
			}
		}
		x = x * 2
	}
	fmt.Println(ar)
	root := BuildTreeWithArr(ar, 1)
	return root
}
func BuildTreeWithString(nodes string) *TreeNode {
	arr := strings.Split(nodes, ",")
	var ar = make([]int, 0)
	var brk bool
	var x, i = 1, 0
	for {
		if brk {
			break
		}
		if i == len(arr) {
			brk = true
			break
		}
		for j := x; j < x+x; j++ {
			if checkParentIsNull(ar, j) {
				ar = append(ar, 0)
				continue
			}
			if arr[i] == "n" {
				i++
				ar = append(ar, 0)
			} else {
				atoi, _ := strconv.Atoi(arr[i])
				ar = append(ar, atoi)
				i++
			}
		}
		x = x * 2
	}
	fmt.Println(ar)
	root := BuildTreeWithArr(ar, 1)
	return root
}
func BuildTreeWithArr(ar []int, start int) *TreeNode {
	if start > len(ar) {
		return nil
	}
	if ar[start-1] == 0 {
		return nil
	}
	//level = 2 * level
	root := &TreeNode{Val: ar[start-1]}
	root.Left = BuildTreeWithArr(ar, 2*start)
	root.Right = BuildTreeWithArr(ar, 2*start+1)
	return root
}
func BuildTreeWithArr1(ar []string, start int) (*TreeNode, int) {
	if start > len(ar) {
		return nil, start
	}
	if ar[start-1] == "null" {
		return nil, start
	}
	//level = 2 * level
	root := &TreeNode{Val: GetIntFromStr(ar[start-1])}
	start++
	root.Left, start = BuildTreeWithArr1(ar, 2*start)
	root.Right, start = BuildTreeWithArr1(ar, 2*start+1)
	return root, start
}

func GetIntFromStr(s string) int {
	atoi, _ := strconv.Atoi(s)
	return atoi
}

func checkParentIsNull(ar []int, j int) bool {
	var isNull bool
	for c := j / 2; c > 0; c = c / 2 {
		if ar[c-1] == 0 {
			ar = append(ar, 0)
			isNull = true
			break
		}
	}
	return isNull
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ret = make([][]int, 0)
	ret = level(root, 0, ret)
	return ret
}

func level(root *TreeNode, l int, ret [][]int) [][]int {
	if root == nil {
		return ret
	}
	if len(ret) < l+1 {
		ret = append(ret, make([]int, 0))
	}
	ret[l] = append(ret[l], root.Val)
	ret = level(root.Left, l+1, ret)
	ret = level(root.Right, l+1, ret)
	return ret
}
