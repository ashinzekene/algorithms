package algorithms

import (
	"strconv"
	"strings"
)

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	switch n {
	case 0:
		return nil
	case 1:
		return lists[0]
	}
	return merge2Lists(mergeKLists(lists[:n/2]), mergeKLists(lists[n/2:]))
}

func merge2Lists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	var head *ListNode
	if l1.Val < l2.Val {
		head = l1
		head.Next = merge2Lists(l1.Next, l2)
	} else {
		head = l2
		head.Next = merge2Lists(l2.Next, l1)
	}
	return head
}

// func mergeKListsSlow(lists []*ListNode) *ListNode {
// 	head := &ListNode{0, nil}
// 	current := head
// 	for len(lists) > 0 {
// 		min, minIndex := -1, -1
// 		for i, l := range lists {
// 			if l != nil {
// 				min, minIndex = l.Val, i
// 				break
// 			}
// 		}
// 		if minIndex == -1 {
// 			return head.Next
// 		}
// 		for i, l := range lists {
// 			if l != nil && l.Val < min {
// 				min = l.Val
// 				minIndex = i
// 			}
// 		}
// 		current.Next = &ListNode{min, nil}
// 		current = current.Next
// 		if lists[minIndex].Next != nil {
// 			lists[minIndex] = lists[minIndex].Next
// 		} else {
// 			lists = append(lists[:minIndex], lists[minIndex+1:]...)
// 		}
// 	}
// 	return head.Next
// }

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	str_node := strconv.Itoa(l.Val)
	current := l
	for current.Next != nil {
		current = current.Next
		str_node += " => " + strconv.Itoa(current.Val)
	}
	return str_node + " => None"
}

func ListNodeFromString(str string, delimiter string) *ListNode {
	if str == "" {
		return nil
	}
	strList := strings.Split(str, delimiter)
	i1, _ := strconv.Atoi(strList[0])
	head := &ListNode{i1, nil}
	current := head
	for _, str := range strList[1:] {
		i, _ := strconv.Atoi(str)
		current.Next = &ListNode{i, nil}
		current = current.Next
	}
	return head
}
