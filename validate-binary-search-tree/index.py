# https://leetcode.com/problems/validate-binary-search-tree/

import math
from typing import Optional

class TreeNode:
    left: 'TreeNode'
    right: 'TreeNode'
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def isValidBST(self, root: Optional[TreeNode]) -> bool:
        return self.isValid(root.left, -math.inf, root.val) and self.isValid(root.right, root.val, math.inf)

    def isValid(self, node: Optional[TreeNode], minVal: int, maxVal: int) -> bool:
        if node == None:
            return True
        if node.val >= maxVal or node.val <= minVal:
            return False

        return self.isValid(node.left, minVal, node.val) and self.isValid(node.right, node.val, maxVal)