package trees

import "github.com/ashinzekene/algorithms/utils"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ========= Create tree from array
func TreeFromList(arr []int) Tree {
	t := Tree{Head: NewTreeNode(arr[0])}
	for _, v := range arr[1:] {
		t.Insert(v)
	}
	return t
}

func TreeFromListLevelOrder(arr []int, nullVal int) *Tree {
	head := createTreeFromListLevelOrder(arr, 0, nullVal)
	return &Tree{
		Head: head,
	}
}

func createTreeFromListLevelOrder(arr []int, i int, nullVal int) *TreeNode {
	// Using 0 to represent nil
	if len(arr) == 0 {
		return nil
	}
	if arr[i] == nullVal {
		return nil
	}
	head := NewTreeNode(arr[i])
	leftIndex := i*2 + 1
	rightIndex := i*2 + 2
	if len(arr) > leftIndex {
		head.Left = createTreeFromListLevelOrder(arr, leftIndex, nullVal)
	}
	if len(arr) > rightIndex {
		head.Right = createTreeFromListLevelOrder(arr, rightIndex, nullVal)
	}
	return head
}

func NewTreeNode(i int) *TreeNode {
	return &TreeNode{Val: i, Left: nil, Right: nil}
}

// ========== Tree Implementation ===============

type Tree struct {
	Head *TreeNode
}

func (t *Tree) Insert(i int) {
	t.Head = insert(i, t.Head)
}

func (t *Tree) Find(i int) *TreeNode {
	return find(i, t.Head)
}

func (t *Tree) Min(i int) *TreeNode {
	node := t.Head
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func (t *Tree) Max(i int) *TreeNode {
	node := t.Head
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func (t *Tree) MaxHeight(i int) int {
	return maxHeight(t.Head)
}

func (t *Tree) MinHeight(i int) int {
	return minHeight(t.Head)
}

func (t *Tree) Remove(i int) {
	t.Head = deleteNode(i, t.Head)
}

func (t *Tree) PreOrder() []int {
	result := make([]int, 0)
	preOrder(t.Head, &result)
	return result
}

func (t *Tree) InOrder() []int {
	result := make([]int, 0)
	inOrder(t.Head, &result)
	return result
}

func (t *Tree) PostOrder() []int {
	result := make([]int, 0)
	postOrder(t.Head, &result)
	return result
}

func (t *Tree) LevelOrder() []int {
	result := make([]int, 0)
	node := t.Head
	queue := []*TreeNode{node}
	for len(queue) > 0 {
		node = queue[0]
		queue = queue[1:]
		result = append(result, node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return result
}

// ========= Utility functions

func insert(i int, node *TreeNode) *TreeNode {
	if node == nil {
		return NewTreeNode(i)
	} else if node.Val > i {
		node.Left = insert(i, node.Left)
		return node
	} else {
		node.Right = insert(i, node.Right)
		return node
	}
}

func find(i int, node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	} else if node.Val == i {
		return node
	} else if node.Val > i {
		return find(i, node.Left)
	} else {
		return find(i, node.Right)
	}
}

func maxHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := 1 + maxHeight(node.Left)
	right := 1 + maxHeight(node.Right)
	return utils.Max(left, right)
}

func minHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := 1 + maxHeight(node.Left)
	right := 1 + maxHeight(node.Right)
	return utils.Min(left, right)
}

func deleteNode(i int, node *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Val > i {
		node.Right = deleteNode(i, node.Right)
		return node
	} else if node.Val < i {
		node.Left = deleteNode(i, node.Left)
		return node
	} else if node.Val == i {
		if node.Left != nil && node.Right != nil {
			// Get min at right
			minNode := node.Right
			for minNode.Left != nil {
				minNode = minNode.Left
			}
			// delete it
			node.Right = deleteNode(minNode.Val, node.Right)
			// replace current node with min node
			node.Val = minNode.Val
			return node
		} else if node.Right != nil {
			return node.Right
		} else if node.Left != nil {
			return node.Left
		}
	}
	return nil
}

func preOrder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	*result = append(*result, node.Val)
	preOrder(node.Left, result)
	preOrder(node.Right, result)
}

func inOrder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	inOrder(node.Left, result)
	*result = append(*result, node.Val)
	inOrder(node.Right, result)
}

func postOrder(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	postOrder(node.Left, result)
	postOrder(node.Right, result)
	*result = append(*result, node.Val)
}
