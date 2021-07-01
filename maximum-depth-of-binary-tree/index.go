package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils"
	. "github.com/ashinzekene/algorithms/utils/trees"
)

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + Max(maxDepth(root.Left), maxDepth(root.Right))
}
