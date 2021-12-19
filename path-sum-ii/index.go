package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils/trees"
)

func PathSum(root *TreeNode, sum int) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	getPathSum(root, &result, sum, []int{})
	return result
}

func getPathSum(root *TreeNode, result *[][]int, sum int, currentPath []int) {
	if root == nil {
		return
	}
	sum -= root.Val
	currentPath = append(currentPath, root.Val)
	if sum == 0 && isRootLeft(root) {
		path := make([]int, len(currentPath))
		copy(path, currentPath)
		*result = append(*result, path)
		return
	}
	getPathSum(root.Left, result, sum, currentPath)
	getPathSum(root.Right, result, sum, currentPath)
}

func isRootLeft(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}
