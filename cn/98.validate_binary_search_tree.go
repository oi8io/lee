//给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
//
// 有效 二叉搜索树定义如下：
//
//
// 节点的左子树只包含 小于 当前节点的数。
// 节点的右子树只包含 大于 当前节点的数。
// 所有左子树和右子树自身必须也是二叉搜索树。
//
//
//
//
// 示例 1：
//
//
//输入：root = [2,1,3]
//输出：true
//
//
// 示例 2：
//
//
//输入：root = [5,1,4,null,null,3,6]
//输出：false
//解释：根节点的值是 5 ，但是右子节点的值是 4 。
//
//
//
//
// 提示：
//
//
// 树中节点数目范围在[1, 10⁴] 内
// -2³¹ <= Node.val <= 2³¹ - 1
//
// 👍 1597 👎 0

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

func isValidBST(root *TreeNode) bool {
	var values []int
	inOrderBST(root, &values)
	if len(values) == 1 {
		return true
	}
	for i := 1; i < len(values); i++ {
		if values[i] <= values[i-1] {
			return false
		}
	}
	return true
}

func inOrderBST(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	inOrderBST(root.Left, values)
	*values = append(*values, root.Val)
	inOrderBST(root.Right, values)
}

/*
func isValidBST(root *TreeNode) bool {
	if !isValidBST1(root.Left, math.MinInt64, root.Val) {
		return false
	}
	return isValidBST1(root.Right, root.Val, math.MaxInt64)
}

func isValidBST1(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val >= max || root.Val <= min {
		return false
	}
	if !isValidBST1(root.Left, min, root.Val) {
		return false
	}
	if !isValidBST1(root.Right, root.Val, max) {
		return false
	}
	return true
}
*/
//leetcode submit region end(Prohibit modification and deletion)
