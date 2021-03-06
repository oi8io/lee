//给定一个二叉树，找出其最大深度。
//
// 二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//
// 说明: 叶子节点是指没有子节点的节点。
//
// 示例：
//给定二叉树 [3,9,20,null,null,15,7]，
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
//
// 返回它的最大深度 3 。
// 👍 1236 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	depth := maxBinaryTreeDepth(root, 0)
	return depth
}

func maxBinaryTreeDepth(root *TreeNode, depth int) int {
	if root == nil {
		return depth
	}
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	leftDepth := maxBinaryTreeDepth(root.Left, depth+1)
	rightDepth := maxBinaryTreeDepth(root.Right, depth+1)
	return max(leftDepth, rightDepth)
}

//leetcode submit region end(Prohibit modification and deletion)
