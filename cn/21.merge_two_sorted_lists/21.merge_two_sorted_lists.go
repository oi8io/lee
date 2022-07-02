//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//
//
//
// 示例 1：
//
//
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
//
//
// 示例 2：
//
//
//输入：l1 = [], l2 = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：l1 = [], l2 = [0]
//输出：[0]
//
//
//
//
// 提示：
//
//
// 两个链表的节点数目范围是 [0, 50]
// -100 <= Node.val <= 100
// l1 和 l2 均按 非递减顺序 排列
//
// 👍 2501 👎 0

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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var answer = &ListNode{}
	h := answer
	n1, n2 := list1, list2
	for n1 != nil || n2 != nil {
		if n1 == nil {
			h.Next = n2
			break
		}
		if n2 == nil {
			h.Next = n1
			break
		}
		if n1.Val <= n2.Val {
			h.Next = n1
			n1 = n1.Next
		} else {
			h.Next = n2
			n2 = n2.Next
		}
		h = h.Next
	}
	return answer.Next
}

//leetcode submit region end(Prohibit modification and deletion)
