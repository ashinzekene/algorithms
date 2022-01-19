package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils/trees"
)

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	nums := getNums(root, 0)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func getNums(root *TreeNode, i int) []int {
	result := make([]int, 0)
	if root.Left == nil && root.Right == nil {
		return []int{i*10 + root.Val}
	}
	if root.Left != nil {
		leftNums := getNums(root.Left, i*10+root.Val)
		result = append(result, leftNums...)
	}
	if root.Right != nil {
		rightNums := getNums(root.Right, i*10+root.Val)
		result = append(result, rightNums...)
	}
	return result
}
