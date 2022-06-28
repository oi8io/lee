//给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
//
//
//
// 示例 1：
//
//
//输入：root = [3,1,4,null,2], k = 1
//输出：1
//
//
// 示例 2：
//
//
//输入：root = [5,3,6,2,4,null,null,1], k = 3
//输出：3
//
//
//
//
//
//
// 提示：
//
//
// 树中的节点数为 n 。
// 1 <= k <= n <= 10⁴
// 0 <= Node.val <= 10⁴
//
//
//
//
// 进阶：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法？
// 👍 625 👎 0

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
func kthSmallest(root *TreeNode, k int) int {
	smallest1, _ := kthSmallest1(root, 0, k)
	return smallest1
}
func kthSmallest1(root *TreeNode, cnt, k int) (int, int) {
	if root == nil {
		return 0, cnt
	}
	smallestL, iL := kthSmallest1(root.Left, cnt, k)
	if iL == k {
		return smallestL, iL
	}
	iL++
	if iL == k {
		return root.Val, iL
	}
	smallestR, iR := kthSmallest1(root.Right, iL, k)
	return smallestR, iR
}

//leetcode submit region end(Prohibit modification and deletion)
