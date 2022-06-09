//给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（
//一个节点也可以是它自己的祖先）。”
//
//
//
// 示例 1：
//
//
//输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
//输出：3
//解释：节点 5 和节点 1 的最近公共祖先是节点 3 。
//
//
// 示例 2：
//
//
//输入：root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
//输出：5
//解释：节点 5 和节点 4 的最近公共祖先是节点 5 。因为根据定义最近公共祖先节点可以为节点本身。
//
//
// 示例 3：
//
//
//输入：root = [1,2], p = 1, q = 2
//输出：1
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [2, 10⁵] 内。
// -10⁹ <= Node.val <= 10⁹
// 所有 Node.val 互不相同 。
// p != q
// p 和 q 均存在于给定的二叉树中。
//
// 👍 1745 👎 0

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
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ancestor236, _ := lowestCommonAncestor236(root, p, q)
	return ancestor236
}
func lowestCommonAncestor236(root, p, q *TreeNode) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	}
	if _, fp := foundTreeNode236(root.Left, p); fp { // 都在左边找到就肯定在左边 继续往下找
		if _, fq := foundTreeNode236(root.Left, q); fq {
			left, b := lowestCommonAncestor236(root.Left, p, q)
			if b {
				return left, true
			}
		}
	}
	if _, rp := foundTreeNode236(root.Right, p); rp {
		if _, rq := foundTreeNode236(root.Right, q); rq {
			left, b := lowestCommonAncestor236(root.Right, p, q)
			if b {
				return left, true
			}
		}
	}
	return root, true
}
func foundTreeNode236(root, n *TreeNode) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	}
	if root == n {
		return root, true
	}
	left, b := foundTreeNode236(root.Left, n)
	if b {
		return left, true
	}
	right, b := foundTreeNode236(root.Right, n)
	if b {
		return right, true
	}
	return root, false
}

//leetcode submit region end(Prohibit modification and deletion)
