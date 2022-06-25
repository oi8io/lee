//给你一棵二叉树的根节点 root ，返回其节点值的 后序遍历 。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[3,2,1]
//
//
// 示例 2：
//
//
//输入：root = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：root = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 树中节点的数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
// 👍 835 👎 0

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
func postorderTraversal(root *TreeNode) []int {
	return postorderTraversal2(root)
	return postorderTraversal1(root)
}

// stack
// 左右头
// 1. 弹出就打印，追加到answer的后面
// 2. 如果有左压入左
// 3. 如果有右压入右
func postorderTraversal2(root *TreeNode) []int {
	var s1 = make([]*TreeNode, 0)
	var s2 = make([]*TreeNode, 0)
	var answer = make([]int, 0)
	head := root
	if head != nil {
		s1 = append(s1, head)
	}
	for len(s1) > 0 {
		n := len(s1) - 1
		head = s1[n]
		s1 = s1[:n]
		s2 = append(s2, head)
		if head.Left != nil {
			s1 = append(s1, head.Left)
		}
		if head.Right != nil {
			s1 = append(s1, head.Right)
		}
	}
	for i := len(s2) - 1; i >= 0; i-- {
		answer = append(answer, s2[i].Val)
	}
	return answer
}

func postorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	answer := make([]int, 0)
	l := postorderTraversal1(root.Left)
	answer = append(answer, l...)
	r := postorderTraversal1(root.Right)
	answer = append(answer, r...)
	answer = append(answer, root.Val)
	return answer
}

//leetcode submit region end(Prohibit modification and deletion)
