package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils/trees"
)

// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/

func flatten(root *TreeNode) {
	root = recurse(root)
}

func recurse(node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	rightNode := node.Right
	if node.Left == nil {
		node.Right = recurse(rightNode)
		return node
	}
	recursed := recurse(node.Left)
	recursedPointer := recursed
	for recursedPointer.Right != nil {
		recursedPointer = recursedPointer.Right
	}
	recursedPointer.Right = recurse(rightNode)
	node.Right = recursed
	node.Left = nil
	return node
}
