//给定一棵二叉树 root，返回所有重复的子树。
//
// 对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
//
// 如果两棵树具有相同的结构和相同的结点值，则它们是重复的。
//
//
//
// 示例 1：
//
//
//
//
//输入：root = [1,2,3,4,null,2,4,null,null,4]
//输出：[[2,4],[4]]
//
// 示例 2：
//
//
//
//
//输入：root = [2,1,1]
//输出：[[1]]
//
// 示例 3：
//
//
//
//
//输入：root = [2,2,2,3,null,3,null]
//输出：[[2,3],[3]]
//
//
//
// 提示：
//
//
// 树中的结点数在[1,10^4]范围内。
// -200 <= Node.val <= 200
//
// 👍 436 👎 0

package cn

import (
	. "github.com/oi8io/lee/cn/common"
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
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var e = make(map[string]int)
	_, ans := serialize(root, e)
	return ans
}

func serialize(root *TreeNode, e map[string]int) (s string, ans []*TreeNode) {
	if root == nil {
		return "#", nil
	}
	l, la := serialize(root.Left, e)
	r, ra := serialize(root.Right, e)
	ans = append(ans, la...)
	ans = append(ans, ra...)
	s = l + "," + r + "," + strconv.Itoa(root.Val)
	if v, ok := e[s]; ok && v == 1 {
		ans = append(ans, root)
	}
	e[s]++
	return s, ans
}

//leetcode submit region end(Prohibit modification and deletion)
