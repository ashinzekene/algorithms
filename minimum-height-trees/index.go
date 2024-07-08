package algorithms

// Topological sort
// Leaf nodes first (one edge)
func findMinHeightTrees(n int, edges [][]int) []int {
	adjList := map[int]map[int]bool{}
	indegrees := make([]int, n)
	result := []int{}
	remainingNodes := n

	for _, edge := range edges {
		edge1 := edge[0]
		edge2 := edge[1]
		if adjList[edge1] == nil {
			adjList[edge1] = map[int]bool{}
		}
		adjList[edge1][edge2] = true
	
		if adjList[edge2] == nil {
			adjList[edge2] = map[int]bool{}
		}
		adjList[edge2][edge1] = true

		indegrees[edge1]++
		indegrees[edge2]++
	}

	for node, edges := range indegrees {
		if edges <= 1 {
			result = append(result, node)
			remainingNodes--
		}
	}

	for remainingNodes >= 1 {
		newResult := []int{}
		for _, node := range result {
			// Only one loop since just one indegree
			for newLeaf := range adjList[node] {
				indegrees[newLeaf]--
				delete(adjList[newLeaf], node) 
				if indegrees[newLeaf] == 1 {
					newResult = append(newResult, newLeaf)
					remainingNodes--
				}
			}
		}

		result = newResult
	}

	return result
}