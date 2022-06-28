//给你二叉树的根结点 root ，请你将它展开为一个单链表：
//
//
// 展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
// 展开后的单链表应该与二叉树 先序遍历 顺序相同。
//
//
//
//
// 示例 1：
//
//
//输入：root = [1,2,5,3,4,null,6]
//输出：[1,null,2,null,3,null,4,null,5,null,6]
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
//输入：root = [0]
//输出：[0]
//
//
//
//
// 提示：
//
//
// 树中结点数在范围 [0, 2000] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？
// 👍 1215 👎 0

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
func flatten(root *TreeNode) {
	flatten2(root)
	return
	nodes := flatten1(root)
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Right = nodes[i+1]
	}
}

func flatten2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := flatten2(root.Left)
	right := flatten2(root.Right)
	root.Left = nil
	root.Right = left
	head := GetRight(root)
	head.Right = right
	return root
}
func GetRight(root *TreeNode) *TreeNode {
	head := root
	for head.Right != nil {
		head = head.Right
	}
	return head
}

func flatten1(root *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}
	left := flatten1(root.Left)
	right := flatten1(root.Right)
	root.Left = nil
	var ret = []*TreeNode{root}
	ret = append(ret, left...)
	ret = append(ret, right...)
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
