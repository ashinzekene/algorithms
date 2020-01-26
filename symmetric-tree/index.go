package algorithms

import "github.com/ashinzekene/algorithms/utils/trees"

func isSymmetric(root *trees.TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left, root.Right)
}

func isMirror(left, right *trees.TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return isMirror(left.Left, right.Right) && isMirror(right.Left, left.Right)
}

func isSymmetric2(root *trees.TreeNode) bool {
	queue := make([]*trees.TreeNode, 0)
	if root == nil {
		return true
	}
	queue = append(queue, root.Left, root.Right)
	for len(queue) > 0 {
		l := queue[0]
		r := queue[1]
		queue = queue[2:]
		if l == nil && r == nil {
			continue
		}
		if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}
		queue = append(queue, l.Left, r.Right)
		queue = append(queue, r.Left, l.Right)
	}
	return true
}
