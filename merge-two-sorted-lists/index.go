package algorithms

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	str_node := strconv.Itoa(l.Val)
	currentNode := l
	for currentNode.Next != nil {
		currentNode = currentNode.Next
		str_node += " => " + strconv.Itoa(currentNode.Val)
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
	currentNode := head
	for _, str := range strList[1:] {
		i, _ := strconv.Atoi(str)
		currentNode.Next = &ListNode{i, nil}
		currentNode = currentNode.Next
	}
	return head
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var currentNode *ListNode
	var otherNode *ListNode
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		head = l1
		otherNode = l2
	} else {
		head = l2
		otherNode = l1
	}
	currentNode = head
	for otherNode != nil && currentNode.Next != nil {
		if currentNode.Next.Val > otherNode.Val {
			oldNext := currentNode.Next
			currentNode.Next = otherNode
			otherNode = oldNext
		}
		currentNode = currentNode.Next
	}
	if currentNode.Next == nil && otherNode != nil {
		currentNode.Next = otherNode
	}
	return head
}
