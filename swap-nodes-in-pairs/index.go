package algorithms

import . "github.com/ashinzekene/algorithms/utils/lists"

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: nil}
	dummy.Next = head
	node := dummy
	current := dummy.Next
	for current != nil && current.Next != nil {
		temp := current.Next
		current.Next = temp.Next
		temp.Next = current
		node.Next = temp
		if temp.Next == nil {
			return dummy.Next
		}
		node = node.Next.Next
		current = temp.Next.Next
	}
	return dummy.Next
}

func swapPairs1(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: nil}
	dummy.Next = head
	current := dummy.Next
	if current != nil && current.Next != nil {
		temp := current.Next
		current.Next = temp.Next
		temp.Next = current
		dummy.Next = temp
		dummy.Next.Next.Next = swapPairs(temp.Next.Next)
	}
	return dummy.Next
}
