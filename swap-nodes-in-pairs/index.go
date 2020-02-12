package algorithms

import . "github.com/ashinzekene/algorithms/utils/lists"

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: nil}
	dummy.Next = head
	node1 := dummy
	if head == nil || head.Next == nil {
		return head
	}
	node2 := head.Next
	for node2 != nil {
		temp := node1.Next
		node1.Next = node2
		temp.Next = node2.Next
		node2.Next = temp

		if temp.Next != nil {
			node2 = temp.Next.Next
		} else {
			return dummy.Next
		}
		node1 = node1.Next.Next
	}

	return dummy.Next
}
