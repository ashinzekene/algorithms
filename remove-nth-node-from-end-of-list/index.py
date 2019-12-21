class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

    def __str__(self):
        str_node = str(self.val)
        current = self
        while current.next:
            current = current.next
            str_node += " => " + str(current.val)
        return str_node + " => None"

def createNode(val: str) -> ListNode:
    head = ListNode(val[0])
    current = head
    for i in val[1:]:
        current.next = ListNode(i)
        current = current.next
    return head

def removeNthFromEnd(head: ListNode, n: int):
        currentNode = head
        nodeCount = 1
        while currentNode.next:
            currentNode = currentNode.next
            nodeCount+=1
        removePoint = nodeCount - n
        if removePoint == 0:
            return head.next
        currentNode = head
        removePoint-=1
        while currentNode.next and removePoint > 0:
            removePoint-=1
            currentNode = currentNode.next
        currentNode.next = currentNode.next.next
        return head

def removeNthFromEnd1(head: ListNode, n: int):
        dummy = ListNode(0)
        dummy.next = head
        first_pointer = dummy
        second_pointer = dummy
        for _ in range(n+1):
            first_pointer = first_pointer.next
        while first_pointer:
            first_pointer = first_pointer.next
            second_pointer = second_pointer.next
        # print(second_pointer.next.next)
        second_pointer.next = second_pointer.next.next
        return dummy.next