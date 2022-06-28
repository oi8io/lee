//给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 8
//
//
//
// 👍 1208 👎 0

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

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

func generateTrees(n int) []*TreeNode {
	answers := tryGenerateTrees(1, n)
	return answers
}

func tryGenerateTrees(start, end int) []*TreeNode {
	if start > end {
		return nil
	}
	if start == end {
		return []*TreeNode{{Val: start}}
	}
	var answer = make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		lefts := tryGenerateTrees(start, i-1)
		rights := tryGenerateTrees(i+1, end)
		if lefts == nil {
			for y := 0; y < len(rights); y++ {
				root := &TreeNode{Val: i}
				root.Right = rights[y]
				answer = append(answer, root)
			}
		} else if rights == nil {
			for x := 0; x < len(lefts); x++ {
				root := &TreeNode{Val: i}
				root.Left = lefts[x]
				answer = append(answer, root)
			}
		} else {
			for x := 0; x < len(lefts); x++ {
				for y := 0; y < len(rights); y++ {
					root := &TreeNode{Val: i}
					root.Left = lefts[x]
					root.Right = rights[y]
					answer = append(answer, root)
				}
			}
		}

	}
	return answer
}

//leetcode submit region end(Prohibit modification and deletion)
