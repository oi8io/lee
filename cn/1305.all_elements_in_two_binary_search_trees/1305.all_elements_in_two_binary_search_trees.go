//给你 root1 和 root2 这两棵二叉搜索树。请你返回一个列表，其中包含 两棵树 中的所有整数并按 升序 排序。.
//
//
//
// 示例 1：
//
//
//
//
//输入：root1 = [2,1,4], root2 = [1,0,3]
//输出：[0,1,1,2,3,4]
//
//
// 示例 2：
//
//
//
//
//输入：root1 = [1,null,8], root2 = [8,1]
//输出：[1,1,8,8]
//
//
//
//
// 提示：
//
//
// 每棵树的节点数在 [0, 5000] 范围内
// -10⁵ <= Node.val <= 10⁵
//
// 👍 160 👎 0

package cn

import . "github.com/oi8io/lee/cn/common"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	a1, a2 := inOrder(root1), inOrder(root2)
	var i, j = 0, 0
	var answer = make([]int, 0)
	for {
		if i == len(a1) {
			answer = append(answer, a2[j:]...)
			break
		}
		if j == len(a2) {
			answer = append(answer, a1[i:]...)
			break
		}
		if a1[i] <= a2[j] {
			answer = append(answer, a1[i])
			i++
		} else {
			answer = append(answer, a2[j])
			j++
		}
	}
	return answer
}

func inOrder(root *TreeNode) (res []int) {
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return
}

//leetcode submit region end(Prohibit modification and deletion)
