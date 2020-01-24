package algorithms

import "github.com/ashinzekene/algorithms/utils/lists"

func removeNthFromEnd(head *lists.ListNode, n int) *lists.ListNode {
	dummy := lists.ListNode{0, nil}
	dummy.Next = head
	firstPointer := &dummy
	secondPointer := &dummy
	for i := 0; i < n+1; i++ {
		firstPointer = firstPointer.Next
	}
	for firstPointer != nil {
		firstPointer = firstPointer.Next
		secondPointer = secondPointer.Next
	}
	secondPointer.Next = secondPointer.Next.Next
	return dummy.Next
}
