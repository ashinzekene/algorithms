package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils/trees"
)

func pathSum(root *TreeNode, targetSum int) int {
	return getPathSum(root, targetSum, 0, map[int]int{0: 1})
}

func getPathSum(node *TreeNode, targetSum, currentSum int, numsMap map[int]int) int {
	result := 0
	if node == nil {
		return result
	}
	currentSum += node.Val
	result += numsMap[currentSum-targetSum]
	numsMap[currentSum]++
	result += getPathSum(node.Left, targetSum, currentSum, numsMap)
	result += getPathSum(node.Right, targetSum, currentSum, numsMap)
	numsMap[currentSum]--
	return result
}
