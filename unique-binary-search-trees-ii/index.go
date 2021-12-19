package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils/trees"
)

func generateTrees(n int) []*TreeNode {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}
	return generate(nums)
}

func generate(nums []int) []*TreeNode {
	if len(nums) == 0 {
		return []*TreeNode{nil}
	}
	if len(nums) == 1 {
		return []*TreeNode{
			NewTreeNode(nums[0]),
		}
	}
	trees := make([]*TreeNode, 0)
	for i, num := range nums {
		possibleLeftTrees := generate(nums[:i])
		possibleRightTrees := generate(nums[i+1:])
		for _, leftTree := range possibleLeftTrees {
			for _, rightTree := range possibleRightTrees {
				trees = append(trees, &TreeNode{
					Val:   num,
					Left:  leftTree,
					Right: rightTree,
				})
			}
		}
	}
	return trees
}
