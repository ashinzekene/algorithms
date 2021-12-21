package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils"
	. "github.com/ashinzekene/algorithms/utils/trees"
)

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(getPossibility(root)...)
}

func getPossibility(root *TreeNode) []int {
	if root == nil {
		return []int{0, 0}
	}
	pRight := getPossibility(root.Right)
	pLeft := getPossibility(root.Left)
	return []int{
		Max(pLeft...) + Max(pRight...),
		pLeft[0] + pRight[0] + root.Val,
	}
}
