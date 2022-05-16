//设计一个算法，找出二叉搜索树中指定节点的“下一个”节点（也即中序后继）。
//
// 如果指定节点没有对应的“下一个”节点，则返回null。
//
// 示例 1:
//
// 输入: root = [2,1,3], p = 1
//
//  2
// / \
//1   3
//
//输出: 2
//
// 示例 2:
//
// 输入: root = [5,3,6,2,4,null,null,1], p = 6
//
//      5
//     / \
//    3   6
//   / \
//  2   4
// /
//1
//
//输出: null
// 👍 120 👎 0

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
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	return inorderSuccessor1(root, p)
}

func inorderSuccessor1(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	successor1 := inorderSuccessor1(root.Left, p)
	if successor1 != nil {
		return successor1
	}
	if root.Val > p.Val {
		return root
	}
	successor2 := inorderSuccessor1(root.Right, p)
	return successor2
}

//leetcode submit region end(Prohibit modification and deletion)
