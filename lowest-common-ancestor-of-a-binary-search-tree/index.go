package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils/trees"
)

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	rootVal := root.Val
	qVal := p.Val
	aVal := q.Val
	if rootVal >= qVal && rootVal <= aVal || rootVal <= qVal && rootVal >= aVal {
		return root
	} else if rootVal > qVal && rootVal > aVal {
		return LowestCommonAncestor(root.Left, p, q)
	} else {
		return LowestCommonAncestor(root.Right, p, q)
	}
}
