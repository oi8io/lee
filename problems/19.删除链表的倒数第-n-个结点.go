/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 *
 * https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (43.80%)
 * Likes:    1954
 * Dislikes: 0
 * Total Accepted:    739.8K
 * Total Submissions: 1.7M
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,3,4,5], n = 2
 * 输出：[1,2,3,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [1], n = 1
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：head = [1,2], n = 1
 * 输出：[1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 链表中结点的数目为 sz
 * 1 <= sz <= 30
 * 0 <= Node.val <= 100
 * 1 <= n <= sz
 *
 *
 *
 *
 * 进阶：你能尝试使用一趟扫描实现吗？
 *
 */

package problems

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var next = head
	var l int
	for {
		if next == nil {
			break
		}
		next = next.Next
		l++
	}
	if l == n {
		return head.Next
	}
	del := head
	for i := 0; i < l-1-n; i++ {
		del = del.Next
	}
	del.Next = del.Next.Next
	return head
}
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	var next = head
	var all []*ListNode
	for {
		if next == nil {
			break
		}
		all = append(all, next)
		next = next.Next
	}
	if len(all) == n {
		return head.Next
	}
	all[len(all)-1-n].Next = all[len(all)-1-n].Next.Next
	return head
}

// @lc code=end
