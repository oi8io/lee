/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 *
 * https://leetcode-cn.com/problems/remove-linked-list-elements/description/
 *
 * algorithms
 * Easy (53.19%)
 * Likes:    862
 * Dislikes: 0
 * Total Accepted:    333.1K
 * Total Submissions: 625.6K
 * Testcase Example:  '[1,2,6,3,4,5,6]\n6'
 *
 * 给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
 *
 *
 * 示例 1：
 *
 *
 * 输入：head = [1,2,6,3,4,5,6], val = 6
 * 输出：[1,2,3,4,5]
 *
 *
 * 示例 2：
 *
 *
 * 输入：head = [], val = 1
 * 输出：[]
 *
 *
 * 示例 3：
 *
 *
 * 输入：head = [7,7,7,7], val = 7
 * 输出：[]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 列表中的节点数目在范围 [0, 10^4] 内
 * 1
 * 0
 *
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
//1、 在头，2、在尾部，3、仅仅包含改元素
//
func removeElements(head *ListNode, val int) *ListNode {
	pre := head
	var list *ListNode
	for {
		if pre == nil {
			break
		}
		// 头部情况
		if pre.Val == val {
			pre = pre.Next
			continue
		}
		// 如果头部去重之后还有剩余则设置为head
		if list == nil {
			list = pre
		}
		if pre.Next == nil {
			break
		}
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next
		} else {
			pre = pre.Next
		}
	}
	return list
}

// @lc code=end
