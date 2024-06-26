from typing import Optional

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    small: TreeNode = None
    large: TreeNode = None
    prev: TreeNode = None
    def recoverTree(self, root: Optional[TreeNode]) -> None:
        """
        Do not return anything, modify root in-place instead.
        """
        self.inOrder(root)
        print("__SMall", self.small.val, "Large", self.large.val)
        self.small.val, self.large.val = self.large.val, self.small.val

    def inOrder(self, root: Optional[TreeNode]):
        if root == None:
            return
        self.inOrder(root.left)
        if self.prev and self.prev.val > root.val:
            if (self.large and self.prev.val > self.large.val) or not self.large:
                self.large = self.prev
            if (self.small and root.val < self.small.val) or not self.small:
                self.small = root
        self.prev = root
        self.inOrder(root.right)
        