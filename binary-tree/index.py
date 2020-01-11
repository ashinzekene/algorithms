class Node():
    left = None
    right = None

    def __init__(self, element):
        self.element = element

    def __str__(self):
        return "\t\t{}\n{}\t\t{}".format(self.element, self.left, self.right)
        

class BinaryTree():
    head = None

    def __init__(self, head):
        self.head = Node(head)
    def __str__(self):
        return self.head.__str__()

    def insert(self, element, node=None):
        if not node:
            node = self.head
        if node.element > element:
            if node.left == None:
                node.left = Node(element)
            else:
                self.insert(element, node.left)
        else:
            if node.right == None:
                node.right = Node(element)
            else:
                self.insert(element, node.right)

    def find(self, element):
        node = self.head
        while node:
            if node.element == element:
                return node
            elif node.element > element:
                node = node.left
            else:
                node = node.right
        return None

    def min(self, node=None):
        if not node:
            node = self.head
        while node.left != None:
            node = node.left
        return node

    def max(self, node=None):
        if not node:
            node = self.head
        while node.right != None:
            node = node.right
        return node

    def remove(self, element, node=None):
        if not node:
            node = self.head
        def removeNode(node):
            if element > node.element:
                node.right = removeNode(node.right)
                return node
            elif element < node.element:
                node.left = removeNode(node.left)
                return node
            else:
                if node.left and node.right:
                    minRightElement = self.min(node.right).element
                    node.right = self.remove(minRightElement, node.right)
                    node.element = minRightElement
                    return node
                elif node.left:
                    return node.left
                elif node.right:
                    return node.right
                else:
                    return None
        node = removeNode(node)
        return node

    def maxHeight(self):
        def nodeHeight(node):
            if not node:
                return 0
            left = 1 + nodeHeight(node.left)
            right = 1 + nodeHeight(node.right)
            return max(left, right)
        return nodeHeight(self.head)

    def minHeight(self):
        def nodeHeight(node):
            if not node:
                return 0
            left = 1 + nodeHeight(node.left)
            right = 1 + nodeHeight(node.right)
            return min(left, right)
        return nodeHeight(self.head)

    def preOrder(self):
        def traverse(node):
            result = []
            if not node:
                return result
            result.append(node.element)
            result.extend(traverse(node.left))
            result.extend(traverse(node.right))
            return result

        return traverse(self.head)

    def inOrder(self):
        def traverse(node):
            result = []
            if not node:
                return result
            result.extend(traverse(node.left))
            result.append(node.element)
            result.extend(traverse(node.right))
            return result

        return traverse(self.head)

    def postOrder(self):
        def traverse(node):
            result = []
            if not node:
                return result
            result.extend(traverse(node.left))
            result.extend(traverse(node.right))
            result.append(node.element)
            return result

        return traverse(self.head)

    def levelOrder(self):
        queue = []
        result = []
        node = self.head
        while node:
            result.append(node.element)
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
            node = None
            if len(queue) > 0:
                node = queue.pop(0)

        return result


bt = BinaryTree(25)
nodes = [15,50,10,22,4,12,18,24,35,70,31,44,66,90]
for node in nodes:
    bt.insert(node)
print(bt.inOrder())
print("===")
print(bt.preOrder())
print("===")
print(bt.postOrder())
print("===")
print(bt.levelOrder())
# print(bt.find(7))
# print(bt.find(5))
# print(bt.min())
# print(bt.max())
# bt.remove(2)
# print(bt.find(2))
print(bt)
bt.remove(4)
print("===")
print(bt)