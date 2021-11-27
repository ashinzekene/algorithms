package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils/lists"
)

func Partition(head *ListNode, x int) *ListNode {
	dummyHead := &ListNode{Val: -201, Next: head}
	pivot, _ := getPivot(dummyHead, x)
	if pivot == nil {
		return head
	}

	paritionPoint := pivot.Next
	current := pivot.Next

	for current != nil && current.Next != nil {
		if current.Next.Val < x {
			val := current.Next.Val
			current.Next = current.Next.Next
			pivot.Next = &ListNode{Val: val, Next: paritionPoint}
			pivot = pivot.Next
		} else {
			current = current.Next
		}
	}

	return dummyHead.Next
}

func getPivot(head *ListNode, x int) (*ListNode, int) {
	current := head
	i := 0
	for current != nil {
		if current.Next != nil && current.Next.Val >= x {
			return current, i
		}
		i++
		current = current.Next
	}
	return nil, -1
}
