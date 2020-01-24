package algorithms

import "github.com/ashinzekene/algorithms/utils/lists"

func mergeKLists(lists []*lists.ListNode) *lists.ListNode {
	n := len(lists)
	switch n {
	case 0:
		return nil
	case 1:
		return lists[0]
	}
	return merge2Lists(mergeKLists(lists[:n/2]), mergeKLists(lists[n/2:]))
}

func merge2Lists(l1 *lists.ListNode, l2 *lists.ListNode) *lists.ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var head *lists.ListNode
	if l1.Val < l2.Val {
		head = l1
		head.Next = merge2Lists(l1.Next, l2)
	} else {
		head = l2
		head.Next = merge2Lists(l2.Next, l1)
	}
	return head
}
