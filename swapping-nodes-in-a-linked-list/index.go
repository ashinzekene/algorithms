package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils"
	. "github.com/ashinzekene/algorithms/utils/lists"
)

func swapNodes(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	// making it 0 index based
	count := getListLength(head)
	node1Index := k
	node2Index := count - k + 1
	node1 := &ListNode{Val: 0, Next: nil}
	node2 := &ListNode{Val: 0, Next: nil}
	currentIndex := 0
	current := dummy
	for currentIndex <= Max(node1Index, node2Index) {
		if currentIndex == node1Index {
			node1 = current
		}
		if currentIndex == node2Index {
			node2 = current
		}
		currentIndex++
		current = current.Next
	}
	node1Val := node1.Val
	node1.Val = node2.Val
	node2.Val = node1Val
	return dummy.Next
}

func getListLength(head *ListNode) int {
	node := head
	count := 1
	for node.Next != nil {
		node = node.Next
		count++
	}
	return count
}
