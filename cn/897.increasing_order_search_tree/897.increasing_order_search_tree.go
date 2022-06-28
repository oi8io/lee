//给你一棵二叉搜索树的 root ，请你 按中序遍历 将其重新排列为一棵递增顺序搜索树，使树中最左边的节点成为树的根节点，并且每个节点没有左子节点，只有一个右
//子节点。
//
//
//
// 示例 1：
//
//
//输入：root = [5,3,6,2,4,null,8,1,null,null,null,7,9]
//输出：[1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]
//
//
// 示例 2：
//
//
//输入：root = [5,1,7]
//输出：[1,null,5,null,7]
//
//
//
//
// 提示：
//
//
// 树中节点数的取值范围是 [1, 100]
// 0 <= Node.val <= 1000
//
// 👍 288 👎 0

package cn

import (
	. "github.com/oi8io/lee/cn/common"
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

func increasingBST(root *TreeNode) *TreeNode {
	head := increasingBST2(root)
	return head
}

func increasingBST2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	hl := increasingBST2(root.Left)
	root.Left = nil
	head := setRight(hl, root)
	hr := increasingBST2(root.Right)
	root.Right = hr
	return head
}
func setRight(root *TreeNode, right *TreeNode) *TreeNode {
	if root == nil {
		return right
	}
	head := root
	for head.Right != nil {
		head = head.Right
	}
	head.Right = right
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
