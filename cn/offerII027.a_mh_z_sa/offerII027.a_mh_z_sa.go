//给定一个链表的 头节点 head ，请判断其是否为回文链表。 
//
// 如果一个链表是回文，那么链表节点序列从前往后看和从后往前看是相同的。 
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入: head = [1,2,3,3,2,1]
//输出: true 
//
// 示例 2： 
//
// 
//
// 
//输入: head = [1,2]
//输出: false
// 
//
// 
//
// 提示： 
//
// 
// 链表 L 的长度范围为 [1, 10⁵] 
// 0 <= node.val <= 9 
// 
//
// 
//
// 进阶：能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？ 
//
// 
//
// 注意：本题与主站 234 题相同：https://leetcode-cn.com/problems/palindrome-linked-list/ 
// 👍 62 👎 0


package cn

import . "github.com/oi8io/lee/cn/common"

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	return isPalindrome3(head)
	return isPalindrome2(head)
	return isPalindrome1(head)
}

// 快慢指针，反向链表
func isPalindrome2(head *ListNode) bool {
	if head.Next == nil {
		return true
	}
	slow := head
	fast := head
	for {
		if fast != nil && fast.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next
		} else {
			break
		}
	}
	n1 := slow
	n2 := n1.Next
	n1.Next = nil
	for n2 != nil {
		n3 := n2.Next
		n2.Next = n1
		n1 = n2
		n2 = n3
	}
	s, e := head, n1
	for s != nil && e != nil {
		if s.Val != e.Val {
			return false
		}
		s, e = s.Next, e.Next
	}
	return true
}
func isPalindrome1(head *ListNode) bool {
	arr := make([]int, 0)
	h := head

	for h != nil {
		arr = append(arr, h.Val)
		h = h.Next
	}
	for i := len(arr) - 1; i >= len(arr)/2; i-- {
		if arr[i] != head.Val {
			return false
		}
		head = head.Next
	}
	return true
}
func isPalindrome3(head *ListNode) bool {
	arr := make([]int, 0)
	h := head

	for h != nil {
		arr = append(arr, h.Val)
		h = h.Next
	}
	for i, j := 0, len(arr)-1; i <= len(arr)/2 && j >= len(arr)/2; i, j = i+1, j-1 {
		if arr[j] != arr[i] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
