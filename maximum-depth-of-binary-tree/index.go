package algorithms

import "github.com/ashinzekene/algorithms/utils/trees"

func maxDepth(root *trees.TreeNode) int {
	if root == nil {
		return 0
	}
	left := 1 + maxDepth(root.Left)
	right := 1 + maxDepth(root.Right)
	if left > right {
		return left
	}
	return right
}
