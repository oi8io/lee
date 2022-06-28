//给定一个单链表的头节点 head ，其中的元素 按升序排序 ，将其转换为高度平衡的二叉搜索树。
//
// 本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差不超过 1。
//
//
//
// 示例 1:
//
//
//
//
//输入: head = [-10,-3,0,5,9]
//输出: [0,-3,9,-10,null,5]
//解释: 一个可能的答案是[0，-3,9，-10,null,5]，它表示所示的高度平衡的二叉搜索树。
//
//
// 示例 2:
//
//
//输入: head = []
//输出: []
//
//
//
//
// 提示:
//
//
// head 中的节点数在[0, 2 * 10⁴] 范围内
// -10⁵ <= Node.val <= 10⁵
//
// 👍 719 👎 0

package cn

import (
	. "github.com/oi8io/lee/cn/common"
)

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	var arr = make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	tree := BuildTreeWithArr(arr, 0, len(arr)-1)
	return tree
}
func BuildTreeWithArr(arr []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}
	if start == end {
		return &TreeNode{Val: arr[start]}
	}
	mid := (start + end) / 2
	root := &TreeNode{Val: arr[mid]}
	if mid != start {
		left := BuildTreeWithArr(arr, start, mid-1)
		root.Left = left
	}
	if mid != end {
		right := BuildTreeWithArr(arr, mid+1, end)
		root.Right = right
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
