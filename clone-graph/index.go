package algorithms

/**
 * Definition for a Node.
 */

type Node struct {
    Val int
    Neighbors []*Node
}
 
 func cloneGraph(node *Node) *Node {
	nodes:= map[int]*Node{}
	return cloneNodes(node, nodes)
}

func cloneNodes(node *Node, nodes map[int]*Node) *Node {
	if node == nil {
			return nil
	}
	if nodes[node.Val] != nil {
			return nodes[node.Val]
	}
	newNode := Node{
			Val: node.Val,
			Neighbors: []*Node{},
	}
	nodes[node.Val] = &newNode
	for _, child := range node.Neighbors {
			newNode.Neighbors = append(newNode.Neighbors, cloneNodes(child, nodes))
	}
	return &newNode
}