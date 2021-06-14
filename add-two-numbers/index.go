package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils/lists"
)

func getSumRemainder(l1Val int, l2Val int, remainder int) (int, int) {
	sum := l1Val + l2Val + remainder
	remainder = sum / 10
	sum = sum % 10
	return sum, remainder
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result = &ListNode{Val: 0, Next: nil}
	resultPointer := result
	remainder := 0

	for l1 != nil && l2 != nil {
		var sum int
		sum, remainder = getSumRemainder(l1.Val, l2.Val, remainder)
		resultPointer.Next = &ListNode{Val: sum, Next: nil}
		resultPointer = resultPointer.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		var sum int
		sum, remainder = getSumRemainder(l1.Val, 0, remainder)
		resultPointer.Next = &ListNode{Val: sum, Next: nil}
		resultPointer = resultPointer.Next
		l1 = l1.Next
	}

	for l2 != nil {
		var sum int
		sum, remainder = getSumRemainder(0, l2.Val, remainder)
		resultPointer.Next = &ListNode{Val: sum, Next: nil}
		resultPointer = resultPointer.Next
		l2 = l2.Next
	}

	if remainder > 0 {
		resultPointer.Next = &ListNode{Val: remainder, Next: nil}
	}

	return result.Next
}
