//给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
//
//
//
// 示例 1：
//
//
//输入：root = [1,null,2,3]
//输出：[1,2,3]
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
// 示例 4：
//
//
//输入：root = [1,2]
//输出：[1,2]
//
//
// 示例 5：
//
//
//输入：root = [1,null,2]
//输出：[1,2]
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
// 进阶：递归算法很简单，你可以通过迭代算法完成吗？
// 👍 817 👎 0

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

// 递归

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	return preorderTraversal2(root)
	return preorderTraversal1(root)
}

/**
头 右  左
*/
func preorderTraversal2(root *TreeNode) []int {
	var stack, answer = make([]*TreeNode, 0), make([]int, 0)
	head := root
	if head != nil {
		stack = append(stack, head)
	}
	for {
		if len(stack) == 0 {
			break
		}
		if len(stack) > 0 {
			n := len(stack) - 1
			head = stack[n] //pop
			stack = stack[:n]
			answer = append(answer, head.Val)
			if head.Right != nil {
				stack = append(stack, head.Right)
			}
			if head.Left != nil {
				stack = append(stack, head.Left)
			}
		}
	}
	return answer
}

//递归
func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var answer = make([]int, 0)
	answer = append(answer, root.Val)
	l := preorderTraversal1(root.Left)
	answer = append(answer, l...)
	r := preorderTraversal1(root.Right)
	answer = append(answer, r...)
	return answer
}

//leetcode submit region end(Prohibit modification and deletion)
