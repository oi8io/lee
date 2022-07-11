//给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。
//
// 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
//
//
//
// 示例1：
//
//
//
//
//输入：l1 = [7,2,4,3], l2 = [5,6,4]
//输出：[7,8,0,7]
//
//
// 示例2：
//
//
//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[8,0,7]
//
//
// 示例3：
//
//
//输入：l1 = [0], l2 = [0]
//输出：[0]
//
//
//
//
// 提示：
//
//
// 链表的长度范围为 [1, 100]
// 0 <= node.val <= 9
// 输入数据保证链表代表的数字无前导 0
//
//
//
//
// 进阶：如果输入链表不能翻转该如何解决？
// 👍 535 👎 0

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers1(l1, l2)
}
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	// todo 反转链表
	return nil
}
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	h1, h2 := l1, l2
	if h1.Val == 0 {
		return l2
	}
	if h2.Val == 0 {
		return l1
	}
	var s1, s2 = make([]*ListNode, 0), make([]*ListNode, 0) //不能反转链表
	for h1 != nil {
		s1 = append(s1, h1)
		h1 = h1.Next
	}
	for h2 != nil {
		s2 = append(s2, h2)
		h2 = h2.Next
	}
	carry := 0
	var ans *ListNode
	for len(s1) > 0 || len(s2) > 0 {
		a, b := 0, 0
		if len(s1) > 0 {
			a = s1[len(s1)-1].Val
			s1 = s1[:len(s1)-1]
		}
		if len(s2) > 0 {
			b = s2[len(s2)-1].Val
			s2 = s2[:len(s2)-1]
		}
		node := &ListNode{Val: (a + b + carry) % 10}
		carry = (a + b + carry) / 10
		node.Next = ans
		ans = node
	}
	if carry > 0 {
		node := &ListNode{Val: carry}
		node.Next = ans
		ans = node
	}
	return ans
}

//leetcode submit region end(Prohibit modification and deletion)
