package algorithms

func isBipartite(graph [][]int) bool {
	n := len(graph)
	adjList := make([][]int, n)
	for i, edges := range graph {
			for _, edge := range edges {
					adjList[i] = append(adjList[i], edge)
					adjList[edge] = append(adjList[edge], i)
			}
	}

	groups := make([]int, n)
	for i := range groups {
			if !dfs(i, -1, groups, adjList) {
					return false
			}
	}
	return true
}

func dfs(node, parent int, groups []int, adjList [][]int) bool {
	if groups[node] != 0 {
			return parent == -1 || groups[node] != groups[parent]
	}
	groups[node] = 0
	if parent != -1 && groups[parent] == 0 {
			groups[node] = 1
	}
	for _, child := range adjList[node] {
			if child == parent {
					continue
			}
			if !dfs(child, node, groups, adjList) {
					return false
			}
	}
	return true
}