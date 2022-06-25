//给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,3,2]
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
// 树中节点数目在范围 [0, 100] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？
// 👍 1422 👎 0

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

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func inorderTraversal(root *TreeNode) []int {
	return inorderTraversal2(root)
	return inorderTraversal1(root)
}

/**
左斜线划分.整条左边界入栈
*/
func inorderTraversal2(root *TreeNode) []int {
	var stack = make([]*TreeNode, 0)
	var answer = make([]int, 0)
	head := root
	n := len(stack)
	for head != nil || len(stack) > 0 {
		if head != nil {
			stack = append(stack, head)
			head = head.Left // 不停将左子树放入到队列中
		} else {
			n = len(stack) - 1
			head = stack[n] //弹出最左边子树
			answer = append(answer, head.Val)
			stack = stack[:n]
			head = head.Right //找右节点的左节点
		}
	}
	return answer
}

func inorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var answer = make([]int, 0)

	l := inorderTraversal1(root.Left)
	answer = append(answer, l...)
	answer = append(answer, root.Val)
	r := inorderTraversal1(root.Right)
	answer = append(answer, r...)
	return answer
}

//leetcode submit region end(Prohibit modification and deletion)
