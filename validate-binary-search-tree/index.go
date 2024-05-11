package algorithms

import (
	"math"

	. "github.com/ashinzekene/algorithms/utils/trees"
)

// https://leetcode.com/problems/validate-binary-search-tree/

func isValidBST(root *TreeNode) bool {
	return validateTree(root, int(math.MaxInt32), int(-math.MaxInt32))
}

func validateTree(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return validateTree(root.Left, min, root.Val) && validateTree(root.Right, root.Val, max)
}
