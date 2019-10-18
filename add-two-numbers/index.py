# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        remainder = 0
        v1 = l1
        v2 = l2
        result = ListNode(0)
        current = result
        while v1 or v2:
            v1Val = 0
            v2Val = 0
            if v1:
                v1Val = v1.val
                v1 = v1.next
            if v2:
                v2Val = v2.val
                v2 = v2.next
            intSum = v1Val + v2Val + remainder
            remainder = intSum // 10
            current.next = ListNode(intSum % 10)
            current = current.next
        if remainder > 0:
            current.next = ListNode(remainder)
        return result.next
            
        