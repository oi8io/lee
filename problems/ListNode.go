package problems

import (
	"fmt"
)

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

func (list *ListNode) String() string {
	var ret string
	head := list
	for {
		if head == nil {
			break
		}
		ret += fmt.Sprintf(" %d ", head.Val)
		head = head.Next
	}
	return ret
}
