package lists

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
