package questions

//给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
//
//
//
// 示例 1：
//
//
//输入：head = [1,2,3,4,5], k = 2
//输出：[4,5,1,2,3]
//
//
// 示例 2：
//
//
//输入：head = [0,1,2], k = 4
//输出：[2,0,1]
//
//
//
//
// 提示：
//
//
// 链表中节点的数目在范围 [0, 500] 内
// -100 <= Node.val <= 100
// 0 <= k <= 2 * 10⁹
//
// Related Topics 链表 双指针 👍 743 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func (head *ListNode) length() int {
	var newHead = head
	n := 1 //先计算长度
	for {
		if newHead.Next == nil {
			break
		}
		newHead = newHead.Next
		n++
	}

	return n
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	var newHead = head
	n := 1 //先计算长度
	for {
		if newHead.Next == nil {
			break
		}
		newHead = newHead.Next
		n++
	}
	k = k % n
	if k == 0 { // 如果是整数倍，则无需再移动
		return head
	}
	k = n - k // 开始元素实际位置为  长度-移动次数-1
	//记录切开位置，表示最后一个元素
	split := head
	for i := 0; i < k-1; i++ {
		split = split.Next
	}
	var list = split.Next
	var end = list //找到末尾位置，将末尾的指针指向原起始位置
	for {
		if end == nil || end.Next == nil {
			break
		}
		end = end.Next
	}
	end.Next = head
	split.Next = nil
	return list
}

//leetcode submit region end(Prohibit modification and deletion)
