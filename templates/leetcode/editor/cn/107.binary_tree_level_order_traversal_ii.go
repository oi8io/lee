//给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：[[15,7],[9,20],[3]]
//
//
// 示例 2：
//
//
//输入：root = [1]
//输出：[[1]]
//
//
// 示例 3：
//
//
//输入：root = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 2000] 内
// -1000 <= Node.val <= 1000
//
// 👍 571 👎 0

package problems

import "fmt"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	h := maxTreeDepth(root)
	//fmt.Println(h)
	result := make([][]int, h)
	fmt.Println(len(result))
	//fmt.Println(cap(result))
	result = levelOrderBottom1(root, result, h-1)
	return result
}

func maxTreeDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lh := maxTreeDepth(root.Left)
	rh := maxTreeDepth(root.Right)
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	h := max(lh, rh) + 1
	return h

}

func levelOrderBottom1(root *TreeNode, result [][]int, depth int) [][]int {
	if root == nil {
		return result
	}
	result = levelOrderBottom1(root.Left, result, depth-1)
	result = levelOrderBottom1(root.Right, result, depth-1)
	if result[depth] == nil {
		result[depth] = make([]int, 0)
	}
	result[depth] = append(result[depth], root.Val)
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
