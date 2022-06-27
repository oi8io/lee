//给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[20,9],[15,7]]
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
// -100 <= Node.val <= 100
//
// 👍 656 👎 0

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

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	answer := make([][]int, 0)
	queue := []*TreeNode{root}
	lvm := make(map[*TreeNode]int)
	lvm[root] = 0

	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		lv := lvm[head]
		if len(answer) <= lv || len(answer[lv]) == 0 {
			answer = append(answer, []int{head.Val})
		} else {
			if lv%2 == 0 { //左到右
				answer[lv] = append(answer[lv], head.Val)
			} else { // 从右向左
				answer[lv] = append([]int{head.Val}, answer[lv]...)
			}
		}
		if t := head.Left; t != nil {
			queue = append(queue, t)
			lvm[t] = lv + 1
		}
		if t := head.Right; t != nil {
			queue = append(queue, t)
			lvm[t] = lv + 1
		}
	}
	return answer
}

//leetcode submit region end(Prohibit modification and deletion)
