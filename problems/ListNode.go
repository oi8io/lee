package problems

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildListNode(in []int) *ListNode {
	var list *ListNode
	for i := len(in) - 1; i >= 0; i-- {
		list = &ListNode{
			Val:  in[i],
			Next: list,
		}
	}
	return list
}

func (head *ListNode) String() string {
	var ret string
	var x *ListNode
	for slow, fast := head, head; fast != nil && fast.Next != nil; {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			x = slow
			break
		}
	}
	list := head
	var hasVisited bool
	for {
		if list == nil {
			break
		}
		if list == x {
			if hasVisited == false {
				hasVisited = true
			} else {
				break
			}
		}
		ret += fmt.Sprintf(" %d ", list.Val)
		list = list.Next
	}
	return ret
}
