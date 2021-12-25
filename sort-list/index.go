package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils/lists"
)

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	} else if head.Next == nil {
		return head
	}
	mid := getMid(head)
	left := sortList(head)
	right := sortList(mid)
	return merge(left, right)
}

func merge(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val > l2.Val {
		l2.Next = merge(l2.Next, l1)
		return l2
	} else {
		l1.Next = merge(l1.Next, l2)
		return l1
	}
}

func getMid(head *ListNode) *ListNode {
	mid := &ListNode{Val: 0, Next: head}
	for head != nil && head.Next != nil {
		mid = mid.Next
		head = head.Next.Next
	}
	newMid := mid.Next
	mid.Next = nil
	return newMid
}
