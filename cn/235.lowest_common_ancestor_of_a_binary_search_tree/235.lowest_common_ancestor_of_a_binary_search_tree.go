//给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
//
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（
//一个节点也可以是它自己的祖先）。”
//
// 例如，给定如下二叉搜索树: root = [6,2,8,0,4,7,9,null,null,3,5]
//
//
//
//
//
// 示例 1:
//
// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
//输出: 6
//解释: 节点 2 和节点 8 的最近公共祖先是 6。
//
//
// 示例 2:
//
// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
//输出: 2
//解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。
//
//
//
// 说明:
//
//
// 所有节点的值都是唯一的。
// p、q 为不同节点且均存在于给定的二叉搜索树中。
//
// 👍 833 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor235main(root, p, q *TreeNode) *TreeNode {
	ancestor235, _ := lowestCommonAncestor235(root, p, q)
	return ancestor235
}
func lowestCommonAncestor235(root, p, q *TreeNode) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	}
	_, fp := foundTreeNode(root.Left, p)
	_, fq := foundTreeNode(root.Left, q)
	if fq && fp { // 都在左边找到就肯定在左边 继续往下找
		left, b := lowestCommonAncestor235(root.Left, p, q)
		if b {
			return left, true
		}
	}
	_, rp := foundTreeNode(root.Right, p)
	_, rq := foundTreeNode(root.Right, q)
	if rp && rq { // 都在右边找到就肯定在右边 继续往下找
		left, b := lowestCommonAncestor235(root.Right, p, q)
		if b {
			return left, true
		}
	}
	return root, true
}
func foundTreeNode(root, n *TreeNode) (*TreeNode, bool) {
	if root == nil {
		return nil, false
	}
	if root == n {
		return root, true
	}
	left, b := foundTreeNode(root.Left, n)
	if b {
		return left, true
	}
	right, b := foundTreeNode(root.Right, n)
	if b {
		return right, true
	}
	return root, false
}

//leetcode submit region end(Prohibit modification and deletion)
