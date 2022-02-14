package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var cur, next *ListNode = nil, head
	for next != nil {
		cur, next.Next, next = next, cur, next.Next
	}
	return cur
}
