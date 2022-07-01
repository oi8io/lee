//给定单个链表的头 head ，使用 插入排序 对链表进行排序，并返回 排序后链表的头 。
//
// 插入排序 算法的步骤:
//
//
// 插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
// 每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
// 重复直到所有输入数据插入完为止。
//
//
// 下面是插入排序算法的一个图形示例。部分排序的列表(黑色)最初只包含列表中的第一个元素。每次迭代时，从输入数据中删除一个元素(红色)，并就地插入已排序的列表
//中。
//
// 对链表进行插入排序。
//
//
//
//
//
// 示例 1：
//
//
//
//
//输入: head = [4,2,1,3]
//输出: [1,2,3,4]
//
// 示例 2：
//
//
//
//
//输入: head = [-1,5,3,4,0]
//输出: [-1,0,3,4,5]
//
//
//
// 提示：
//
//
//
//
// 列表中的节点数在 [1, 5000]范围内
// -5000 <= Node.val <= 5000
//
// 👍 526 👎 0

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
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n1, n2 := head, head.Next
	//sh, se := head, head
	var sh, se *ListNode
	for n1 != nil {
		switch {
		case sh == nil:
			sh = n1
			se = sh
			sh.Next = nil
		case n1.Val >= se.Val:
			se.Next = n1
			n1.Next = nil
			se = n1
		case n1.Val <= sh.Val:
			n1.Next = sh
			sh = n1
		default:
			Insert(sh, n1)
		}
		n1 = n2
		if n1 != nil {
			n2 = n1.Next
		}
	}
	return sh
}

func Insert(head *ListNode, n *ListNode) {
	n1, n2 := head, head.Next
	for n1 != nil {
		if n2.Val >= n.Val {
			break
		}
		n1 = n2
		if n1 != nil {
			n2 = n1.Next
		}
	}
	n.Next = n2
	n1.Next = n
}

//leetcode submit region end(Prohibit modification and deletion)
