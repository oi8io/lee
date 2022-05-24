//如果二叉树每个节点都具有相同的值，那么该二叉树就是单值二叉树。
//
// 只有给定的树是单值二叉树时，才返回 true；否则返回 false。
//
//
//
// 示例 1：
//
//
//
// 输入：[1,1,1,1,1,null,1]
//输出：true
//
//
// 示例 2：
//
//
//
// 输入：[2,2,2,5,2]
//输出：false
//
//
//
//
// 提示：
//
//
// 给定树的节点数范围是 [1, 100]。
// 每个节点的值都是整数，范围为 [0, 99] 。
//
// 👍 116 👎 0

package problems

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
	return isUnivalTree1(root, root.Val)
}

func isUnivalTree1(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}
	if root.Val != val {
		return false
	}
	if tree1 := isUnivalTree1(root.Left, val); !tree1 {
		return false
	}
	if tree1 := isUnivalTree1(root.Right, val); !tree1 {
		return false
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
