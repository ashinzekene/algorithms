package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils/lists"
)

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := &ListNode{Val: 0, Next: head}
	current := dummyHead
	for i := 1; i < left; i++ {
		current = current.Next
	}
	toReverse := current.Next
	toReverseCurrent := toReverse.Next
	toReverse.Next = nil
	toReverseTailPointer := toReverse
	for i := left; i < right; i++ {
		oldToRevere := toReverse
		toReverse = toReverseCurrent
		toReverseCurrent = toReverseCurrent.Next
		toReverse.Next = oldToRevere
	}
	toReverseTailPointer.Next = toReverseCurrent
	current.Next = toReverse
	return dummyHead.Next
}
