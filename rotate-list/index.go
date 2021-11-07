package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils/lists"
)

func RotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	length := getListLength(head)
	k %= length
	point := length - k
	if point == length {
		return head
	}
	current := head
	newHead := &ListNode{
		Val:  0,
		Next: current,
	}

	for i := 1; i < point; i++ {
		current = current.Next
	}
	newStart := current.Next
	newEnd := newHead.Next
	current.Next = nil

	current = newStart
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newEnd
	return newStart
}

func getListLength(head *ListNode) int {
	result := 0
	current := head
	for current != nil {
		current = current.Next
		result++
	}
	return result
}
