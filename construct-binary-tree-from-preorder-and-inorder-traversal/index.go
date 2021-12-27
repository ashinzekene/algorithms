package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils/trees"
)

func buildTree(preorder []int, inorder []int) *TreeNode {
	return iBuildTree(&preorder, inorder)
}

func iBuildTree(preorder *[]int, inorder []int) *TreeNode {
	if len(*preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	head := &TreeNode{Val: (*preorder)[0]}
	*preorder = (*preorder)[1:]
	left, right := divideTree(inorder, head.Val)
	head.Left = iBuildTree(preorder, left)
	head.Right = iBuildTree(preorder, right)
	return head
}

func divideTree(inorder []int, headVal int) ([]int, []int) {
	headIndex := 0
	for i, v := range inorder {
		if v == headVal {
			headIndex = i
			break
		}
	}
	return inorder[0:headIndex], inorder[headIndex+1:]
}
