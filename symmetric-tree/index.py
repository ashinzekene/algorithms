from typing import Optional

class TreeNode:
    left: 'TreeNode'
    right: 'TreeNode'
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

# Recursive approach
class Solution:
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        return self.isSame(root.left, root.right)

    def isSame(self, left: Optional[TreeNode], right: Optional[TreeNode]) -> bool:
        if left == None and right == None:
            return True
        if left == None or right == None:
            return False
        if left.val != right.val:
            return False
        return self.isSame(left.left, right.right) and self.isSame(left.right, right.left)


# Iterative approach
class Solution2:
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        queue: list[TreeNode] = []
        # Ideally we always want at least 2 items to compare with
        while len(queue) > 1:
            # Dequeue two items from the queue
            left = queue[0]
            right = queue[1]
            queue = queue[2:]

            # check if they are symmetric
            if left == None and right == None:
                continue
            if left == None or right == None:
                return False
            if left.val != right.val:
                return False
            
            # enqueue left of right and right of left to compare and vice versa
            # like we did in the recursive approach
            queue.extend([left.right, right.left])
            queue.extend([left.left, right.right])
            
        return True
            