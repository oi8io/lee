//给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。请你反转从位置 left 到位置 right 的链
//表节点，返回 反转后的链表 。
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], left = 2, right = 4
//输出：[1,4,3,2,5]
//
//
// 示例 2：
//
//
//输入：head = [5], left = 1, right = 1
//输出：[5]
//
//
//
//
// 提示：
//
//
// 链表中节点数目为 n
// 1 <= n <= 500
// -500 <= Node.val <= 500
// 1 <= left <= right <= n
//
//
//
//
// 进阶： 你可以使用一趟扫描完成反转吗？
// 👍 1329 👎 0

package cn

import (
	"fmt"
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
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	n1, n2 := head, head.Next
	if n2 == nil || left == right {
		return head
	}
	stack := make([]*ListNode, 0)
	cnt := 1
	var bf, af *ListNode
	var s, e *ListNode
	var started bool

	for {
		//fmt.Println(cnt, n1)
		if cnt == left {
			started = true
			s = n1
		}
		if cnt == right+1 {
			af = n1
			e = n1
			break
		}
		if !started {
			bf = n1
		} else {
			stack = append(stack, n1)
			//n1.Next = n2.Next
			//n2.Next = n1
		}
		n1 = n1.Next
		if n1 != nil {
			n2 = n1.Next
		}
		cnt++
	}
	_, _ = s, e
	for i := len(stack) - 1; i >= 1; i-- {
		stack[i].Next = stack[i-1]
	}
	if bf == nil {
		head = stack[len(stack)-1]
	} else {
		bf.Next = stack[len(stack)-1]
	}
	//fmt.Println(stack[len(stack)-1])
	stack[0].Next = af
	return head
}
func reverseBetween1(head *ListNode, left int, right int) *ListNode {
	return nil
	n1, n2 := head, head.Next
	if n2 == nil || left == right {
		return head
	}

	var bf, af, s, e *ListNode
	// 分别记录left的前一个位置。right的后一个位置
	fmt.Println(af, bf, s, e)
	var started bool
	for n1.Val != right {
		if n1.Val == left {
			started = true
		}
		if n1.Val == right {
			return nil
		}
		if !started {
			bf = n1
			n1 = n1.Next
			n2 = n1.Next
		} else {
			n1 = n2.Next
		}
	}

	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
