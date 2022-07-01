//给定一个单链表 L 的头节点 head ，单链表 L 表示为：
//
//
//L0 → L1 → … → Ln - 1 → Ln
//
//
// 请将其重新排列后变为：
//
//
//L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
//
// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//
//
//
// 示例 1：
//
//
//
//
//输入：head = [1,2,3,4]
//输出：[1,4,2,3]
//
// 示例 2：
//
//
//
//
//输入：head = [1,2,3,4,5]
//输出：[1,5,2,4,3]
//
//
//
// 提示：
//
//
// 链表的长度范围为 [1, 5 * 10⁴]
// 1 <= node.val <= 1000
//
// 👍 951 👎 0

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
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	reorderList2(head)
	return
	reorderList1(head)
}
func reorderList2(head *ListNode) {
	h := head
	var cnt = 0
	var q = make([]int, 0)
	for h != nil {
		q = append(q, h.Val)
		if cnt%2 == 0 {
			h.Val = q[0]
			q = q[1:]
		}
		h = h.Next
		cnt++
	}
	h = head.Next
	//fmt.Println(q)
	for h != nil && len(q) > 0 {
		h.Val = q[len(q)-1]
		q = q[:len(q)-1]
		if h.Next != nil {
			h = h.Next.Next
		}
	}
	//fmt.Println(head)
}
func reorderList1(head *ListNode) {
	if head == nil {
		return
	}
	h := head
	var arr = make([]*ListNode, 0)
	var vals = make([]int, 0)
	for h != nil {
		arr = append(arr, h)
		vals = append(vals, h.Val)
		h = h.Next
	}
	n := len(arr)
	for i := 0; i < n; i = i + 2 {
		arr[i].Val = vals[i/2]
	}
	for i := 1; i < n; i = i + 2 {
		arr[i].Val = vals[n-1-i/2]
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
