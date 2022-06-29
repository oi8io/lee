//给你一个链表数组，每个链表都已经按升序排列。
//
// 请你将所有链表合并到一个升序链表中，返回合并后的链表。
//
//
//
// 示例 1：
//
// 输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
//
//
// 示例 2：
//
// 输入：lists = []
//输出：[]
//
//
// 示例 3：
//
// 输入：lists = [[]]
//输出：[]
//
//
//
//
// 提示：
//
//
// k == lists.length
// 0 <= k <= 10^4
// 0 <= lists[i].length <= 500
// -10^4 <= lists[i][j] <= 10^4
// lists[i] 按 升序 排列
// lists[i].length 的总和不超过 10^4
//
// 👍 2030 👎 0

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
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists)==0 {
		return nil
	}
	for len(lists) > 1 {
		lists = mergeHalfLists(lists)
	}
	return lists[0]
}

func mergeHalfLists(lists []*ListNode) []*ListNode {
	var answer = make([]*ListNode, 0)
	for i := 1; i < len(lists); {
		a := mergeTwoLists(lists[i-1], lists[i])
		answer = append(answer, a)
		i = i + 2
	}
	if len(lists)%2 == 1 {
		answer = append(answer, lists[len(lists)-1])
	}
	return answer
}

func mergeTwoLists(a, b *ListNode) *ListNode {
	var ret = new(ListNode)
	var head = ret
	for {
		if a == nil {
			head.Next = b
			break
		}
		if b == nil {
			head.Next = a
			break
		}
		if a.Val >= b.Val {
			head.Next = b
			b = b.Next
		} else {
			head.Next = a
			a = a.Next
		}
		head = head.Next
	}
	return ret.Next
}

//leetcode submit region end(Prohibit modification and deletion)
