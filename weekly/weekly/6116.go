package weekly

import . "github.com/oi8io/lee/cn/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func evaluateTree(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == 1
	}
	l := evaluateTree(root.Left)
	r := evaluateTree(root.Right)
	if root.Val == 3 {
		return l && r
	}
	return l || r
}
