//给你一个链表，两两交换其中相邻的节点，并返回交换后链表的头节点。你必须在不修改节点内部的值的情况下完成本题（即，只能进行节点交换）。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4]
//输出：[2,1,4,3]
//
//
// 示例 2：
//
//
//输入：head = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 100] 内
// 0 <= Node.val <= 100
//
// 👍 1442 👎 0

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
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	n1, n2 := head, head.Next
	if n2 == nil {
		return head
	}
	head = n2
	nx := n1
	for n2 != nil {
		nx.Next = n2
		n1.Next = n2.Next
		n2.Next = n1
		nx = n1
		n1 = n1.Next
		if n1 == nil {
			break
		}
		n2 = n1.Next
		_ = nx.Next

	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
