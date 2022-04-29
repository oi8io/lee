package problems

import "fmt"

//给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
//
// k 是一个正整数，它的值小于或等于链表的长度。
//
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。
//
// 进阶：
//
//
// 你可以设计一个只使用常数额外空间的算法来解决此问题吗？
// 你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
//
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]
//
//
// 示例 2：
//
//
//输入：head = [1,2,3,4,5], k = 3
//输出：[3,2,1,4,5]
//
//
// 示例 3：
//
//
//输入：head = [1,2,3,4,5], k = 1
//输出：[1,2,3,4,5]
//
//
// 示例 4：
//
//
//输入：head = [1], k = 1
//输出：[1]
//
//
//
//
//
// 提示：
//
//
// 列表中节点的数量在范围 sz 内
// 1 <= sz <= 5000
// 0 <= Node.val <= 1000
// 1 <= k <= sz
//
// Related Topics 递归 链表 👍 1554 👎 0

//leetcode submit region begin(Prohibit modification and deletion)

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildList(in []int) *ListNode {
	var list *ListNode
	for i := len(in) - 1; i >= 0; i-- {
		i2 := in[i]
		if list == nil {
			list = &ListNode{
				Val:  i2,
				Next: nil,
			}
		} else {
			current := &ListNode{
				Val:  i2,
				Next: list,
			}
			list = current
		}
	}
	return list
}

func (l *ListNode) String() string {
	var ret = fmt.Sprintf("%d ", l.Val)
	if l.Next != nil {
		ret += l.Next.String()
	}
	return ret
}

/**
解题思路，采用空间换时间，将指针插入到数组中，根据数组的有序性进行从后往前遍历
边界条件：
1. 如果链表长度的是n的整数倍，则最后的newHeader未空指针，要把最后链表元素的Next指针置为nil
2. 如果链表长度不是n的整数倍，则最后将newHeader追加到已经旋转的链表中
*/
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k == 1 {
		return head
	}
	var rev []*ListNode
	var newHead = head
	for i := 0; i < k; i++ {
		if newHead == nil {
			break
		}
		rev = append(rev, newHead)
		newHead = newHead.Next
	}
	var list, cur *ListNode
	//表示没有被中断过
	if len(rev) == k {
		for i := len(rev) - 1; i >= 0; i-- {
			if list == nil {
				list = rev[i]
				cur = list
			} else {
				cur.Next = rev[i]
				cur = cur.Next
			}
		}
		if newHead == nil {
			cur.Next = nil
		} else {
			cur.Next = reverseKGroup(newHead, k)
		}
	} else {
		list = head
	}
	return list
}

//leetcode submit region end(Prohibit modification and deletion)
