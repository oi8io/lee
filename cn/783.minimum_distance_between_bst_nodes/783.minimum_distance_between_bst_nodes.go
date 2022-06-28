//给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
//
// 差值是一个正数，其数值等于两值之差的绝对值。
//
//
//
//
//
// 示例 1：
//
//
//输入：root = [4,2,6,1,3]
//输出：1
//
//
// 示例 2：
//
//
//输入：root = [1,0,48,null,null,12,49]
//输出：1
//
//
//
//
// 提示：
//
//
// 树中节点的数目范围是 [2, 100]
// 0 <= Node.val <= 10⁵
//
//
//
//
// 注意：本题与 530：https://leetcode-cn.com/problems/minimum-absolute-difference-in-
//bst/ 相同
//
//
// 👍 222 👎 0

package cn

import "math"
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
func minDiffInBST(root *TreeNode) int {
	arr := inOrder(root)
	answer := math.MaxInt
	for i := 1; i < len(arr); i++ {
		v := arr[i] - arr[i-1]
		if v < answer {
			answer = v
		}
	}
	return answer
}

func inOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var answer = make([]int, 0)
	l := inOrder(root.Left)
	answer = append(answer, l...)
	answer = append(answer, root.Val)
	r := inOrder(root.Right)
	answer = append(answer, r...)
	return answer
}

//leetcode submit region end(Prohibit modification and deletion)
