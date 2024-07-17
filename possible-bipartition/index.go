package algorithms

func possibleBipartition(n int, dislikes [][]int) bool {
	adjList := make([][]int, n+1)
	for _, dislike := range dislikes {
			a := dislike[0]
			b := dislike[1]
			adjList[a] = append(adjList[a], b)
			adjList[b] = append(adjList[b], a)
	}

	groups := make([]int, n+1)
	for i := range groups {
			groups[i] = -1
	}
	groups[0] = -2
	for i := 1; i <=n;i++ {
		if !dfs(0, i, groups, adjList) {
			return false
		}
	}
	return true
}

func dfs(parent, node int, groups []int, adjList [][]int) bool {
	if groups[node] != -1 {
			return groups[node] != groups[parent]
	}
	groups[node] = 0
	if groups[parent] == 0 {
			groups[node] = 1
	}
	for _, edge := range adjList[node] {
			if edge == parent {
					continue
			}
			if !dfs(node, edge, groups, adjList) {
					return false
			}
	}
	return true
}