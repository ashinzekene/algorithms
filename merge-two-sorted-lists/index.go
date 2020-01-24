package algorithms

import "github.com/ashinzekene/algorithms/utils/lists"

func mergeTwoLists(l1 *lists.ListNode, l2 *lists.ListNode) *lists.ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var head *lists.ListNode
	if l1.Val < l2.Val {
		head = l1
		head.Next = mergeTwoLists(l1.Next, l2)
	} else {
		head = l2
		head.Next = mergeTwoLists(l2.Next, l1)
	}
	return head
}

// func mergeTwoLists(l1 *lists.ListNode, l2 *lists.ListNode) *lists.ListNode {
// 	var head *lists.ListNode
// 	var currentNode *lists.ListNode
// 	var otherNode *lists.ListNode
// 	if l1 == nil {
// 		return l2
// 	} else if l2 == nil {
// 		return l1
// 	}
// 	if l1.Val < l2.Val {
// 		head = l1
// 		otherNode = l2
// 	} else {
// 		head = l2
// 		otherNode = l1
// 	}
// 	currentNode = head
// 	for otherNode != nil && currentNode.Next != nil {
// 		if currentNode.Next.Val > otherNode.Val {
// 			oldNext := currentNode.Next
// 			currentNode.Next = otherNode
// 			otherNode = oldNext
// 		}
// 		currentNode = currentNode.Next
// 	}
// 	if currentNode.Next == nil && otherNode != nil {
// 		currentNode.Next = otherNode
// 	}
// 	return head
// }
