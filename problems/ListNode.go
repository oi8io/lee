package problems

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
