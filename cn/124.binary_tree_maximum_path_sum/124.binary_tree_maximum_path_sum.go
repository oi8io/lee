//路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不
//一定经过根节点。
//
// 路径和 是路径中各节点值的总和。
//
// 给你一个二叉树的根节点 root ，返回其 最大路径和 。
//
//
//
// 示例 1：
//
//
//输入：root = [1,2,3]
//输出：6
//解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
//
// 示例 2：
//
//
//输入：root = [-10,9,20,null,null,15,7]
//输出：42
//解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42
//
//
//
//
// 提示：
//
//
// 树中节点数目范围是 [1, 3 * 10⁴]
// -1000 <= Node.val <= 1000
//
// 👍 1636 👎 0

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

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func maxPathSum(root *TreeNode) int {
	sum1, _ := maxPathSum1(root)
	return sum1
}

func maxPathSum1(root *TreeNode) (s int, p int) {
	if root.Left == nil && root.Right == nil {
		return root.Val, root.Val
	}
	//当前路径和 左+右 + root
	// maxPath( left,right)
	maxSum, maxPathL, maxPathR := root.Val, root.Val, root.Val
	if root.Left != nil {
		ls, lp := maxPathSum1(root.Left)
		if lp > 0 {
			maxPathL += lp
		}
		maxSum = max(ls, maxSum)
	}
	if root.Right != nil {
		rs, rp := maxPathSum1(root.Right)
		if rp > 0 {
			maxPathR += rp
		}
		maxSum = max(rs, maxSum)
	}
	maxSum = max(maxSum, maxPathL+maxPathR-root.Val)
	return maxSum, max(maxPathL, maxPathR)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
