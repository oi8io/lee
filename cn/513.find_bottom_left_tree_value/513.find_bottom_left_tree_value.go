//给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。
//
// 假设二叉树中至少有一个节点。
//
//
//
// 示例 1:
//
//
//
//
//输入: root = [2,1,3]
//输出: 1
//
//
// 示例 2:
//
//
//
//
//输入: [1,2,3,4,null,5,6,null,null,7]
//输出: 7
//
//
//
//
// 提示:
//
//
// 二叉树的节点个数的范围是 [1,10⁴]
// -2³¹ <= Node.val <= 2³¹ - 1
//
// 👍 292 👎 0

package cn

import (
	. "github.com/oi8io/lee/cn/tree"
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
func findBottomLeftValue(root *TreeNode) int {
	// 分别返回左右子树的值，以及高度
	value1, _ := findBottomLeftValue1(root, 0)
	return value1
}
func findBottomLeftValue1(root *TreeNode, level int) (int, int) {
	if root == nil {
		return -1, 0
	}
	var lv, ll, rv, rl int
	// 左子树为空,val 为本身
	if root.Left == nil {
		//	子节点
		lv, ll = root.Val, level
	} else {
		lv, ll = findBottomLeftValue1(root.Left, level+1)
	}
	// 又子树为空，val 为左子树，或本身 或 父节点
	if root.Right == nil {
		rl = -1
	} else {
		rv, rl = findBottomLeftValue1(root.Right, level+1)
	}
	if ll >= rl {
		return lv, ll
	} else {
		return rv, rl
	}
}

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

//leetcode submit region end(Prohibit modification and deletion)
