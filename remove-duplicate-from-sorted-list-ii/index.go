package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils/lists"
)

func DeleteDuplicates(head *ListNode) *ListNode {
	dummyHead := &ListNode{
		Val:  101, // vals are of the range -100 to 100
		Next: head,
	}

	first := dummyHead
	second := dummyHead.Next
	for second != nil && second.Next != nil {
		if first.Next.Val == second.Next.Val {
			currentVal := second.Next.Val
			for second.Next != nil && second.Next.Val == currentVal {
				second = second.Next
			}
			if second.Next == nil {
				first.Next = nil
				continue
			}
			first.Next = second.Next
			second = second.Next
		} else {
			first = first.Next
			second = second.Next
		}
	}
	return dummyHead.Next
}
