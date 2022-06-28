//给你二叉搜索树的根节点 root ，该树中的 恰好 两个节点的值被错误地交换。请在不改变其结构的情况下，恢复这棵树 。
//
//
//
// 示例 1：
//
//
//输入：root = [1,3,null,null,2]
//输出：[3,1,null,null,2]
//解释：3 不能是 1 的左孩子，因为 3 > 1 。交换 1 和 3 使二叉搜索树有效。
//
//
// 示例 2：
//
//
//输入：root = [3,1,4,null,null,2]
//输出：[2,1,4,null,null,3]
//解释：2 不能在 3 的右子树中，因为 2 < 3 。交换 2 和 3 使二叉搜索树有效。
//
//
//
// 提示：
//
//
// 树上节点的数目在范围 [2, 1000] 内
// -2³¹ <= Node.val <= 2³¹ - 1
//
//
//
//
// 进阶：使用 O(n) 空间复杂度的解法很容易实现。你能想出一个只使用 O(1) 空间的解决方案吗？
// 👍 743 👎 0

package cn

import (
	"fmt"
	. "github.com/oi8io/lee/cn/449.serialize_and_deserialize_bst"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

type info struct {
	bl  *TreeNode // bad left
	br  *TreeNode // bad right
	max *TreeNode
	min *TreeNode
}

func (i *info) String() string {
	return ""
	return fmt.Sprintf("bl [%d] br [%d] min [%d] max [%d]", i.bl.Val, i.br.Val, i.max.Val, i.min.Val)
}

func recoverTree(root *TreeNode) {
	stack := []*TreeNode{}
	var x, y, pred *TreeNode
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if pred != nil && root.Val < pred.Val {
			y = root
			if x == nil {
				x = pred
			} else {
				break
			}
		}
		pred = root
		root = root.Right
	}
	x.Val, y.Val = y.Val, x.Val
}
//递归中序遍历
func recoverTree2(root *TreeNode) {
	nodes := inOrder(root)
	index1, index2 := -1, -1
	for i := 0; i < len(nodes)-1; i++ {
		if nodes[i].Val > nodes[i+1].Val {
			index2 = i + 1
			if index1 == -1 {
				index1 = i
			} else {
				break
			}
		}
	}
	x, y := nodes[index1], nodes[index2]
	x.Val, y.Val = y.Val, x.Val
	//fmt.Println(bad[0].Val, bad[1].Val)
	return
	recoverTree1(root)
}
func inOrder(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	var nodes = make([]*TreeNode, 0)
	left := inOrder(root.Left)
	nodes = append(nodes, left...)
	nodes = append(nodes, root)
	right := inOrder(root.Right)
	nodes = append(nodes, right...)
	return nodes
}

func recoverTree1(root *TreeNode) *info {
	if root == nil {
		return nil
	}
	i := &info{
		bl:  nil,
		br:  nil,
		max: root,
		min: root,
	}
	if root.Left != nil {
		il := recoverTree1(root.Left)
		i, finish := check(il, i)
		if finish {
			return i
		}
		if il.max.Val > root.Val {
			i.bl = il.max
		}
	}
	if root.Right != nil {
		ir := recoverTree1(root.Right)
		i, finish := check(ir, i)
		if finish {
			return i
		}
		if ir.min.Val < root.Val {
			i.br = ir.min
		}
	}
	return i

}
func check(i *info, n *info) (*info, bool) {
	if i == nil {
		return n, false
	}
	n.bl = i.bl
	n.br = i.br
	if n.br != nil || n.bl != nil {
		return n, true
	}
	return n, false
}

//leetcode submit region end(Prohibit modification and deletion)
