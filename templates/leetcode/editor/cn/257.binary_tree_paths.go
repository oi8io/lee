//给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。
//
// 叶子节点 是指没有子节点的节点。
//
//
// 示例 1：
//
//
//输入：root = [1,2,3,null,5]
//输出：["1->2->5","1->3"]
//
//
// 示例 2：
//
//
//输入：root = [1]
//输出：["1"]
//
//
//
//
// 提示：
//
//
// 树中节点的数目在范围 [1, 100] 内
// -100 <= Node.val <= 100
//
// 👍 734 👎 0

package problems

import (
	"strconv"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	var paths = make([]string, 0)
	binaryTreePath(root, "", &paths)
	return paths
}

func binaryTreePath(root *TreeNode, path string, paths *[]string) {
	if root == nil {
		return
	}
	if path == "" {
		path = strconv.Itoa(root.Val)
	} else {
		path += "->" + strconv.Itoa(root.Val)
	}
	if root.Left == nil && root.Right == nil {
		*paths = append(*paths, path)
	}
	binaryTreePath(root.Left, path, paths)
	binaryTreePath(root.Right, path, paths)
}

//leetcode submit region end(Prohibit modification and deletion)
