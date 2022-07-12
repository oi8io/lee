//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//
//
// 示例 2：
//
//
//输入：head = [1], n = 1
//输出：[]
//
//
// 示例 3：
//
//
//输入：head = [1,2], n = 1
//输出：[1]
//
//
//
//
// 提示：
//
//
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
//
//
//
//
// 进阶：你能尝试使用一趟扫描实现吗？
// 👍 2122 👎 0

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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var a = make([]*ListNode,0)
	h:=head
	for h!=nil  {
		a= append(a, h)
		h=h.Next
	}
	if len(a)==1 {
		return nil
	}
	switch n {
	case len(a):
		return a[1]
	case 0:
		a[len(a)-2].Next=nil
		return a[0]
	default:
		a[len(a)-n-1].Next=a[len(a)-n].Next
		return a[0]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
