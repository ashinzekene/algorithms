package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils/trees"
)

func buildTree(inorder []int, postorder []int) *TreeNode {
	return iBuildTree(inorder, &postorder)
}

func iBuildTree(inorder []int, postorder *[]int) *TreeNode {
	if len(inorder) == 0 || len(*postorder) == 0 {
		return nil
	}
	lPost := len(*postorder)
	head := &TreeNode{Val: (*postorder)[lPost-1]}
	*postorder = (*postorder)[:lPost-1]
	left, right := divideTree(inorder, head.Val)
	head.Right = iBuildTree(right, postorder)
	head.Left = iBuildTree(left, postorder)
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
