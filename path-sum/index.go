package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils/trees"
)

func HasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if targetSum == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	return HasPathSum(root.Right, targetSum) || HasPathSum(root.Left, targetSum)
}
