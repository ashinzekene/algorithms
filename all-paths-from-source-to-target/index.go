package algorithms

func allPathsSourceTarget(graph [][]int) [][]int {
	result := make([][]int, 0)
	dfs(0, graph, &result, []int{0})
	return result
}

func dfs(curr int, graph [][]int, result *[][]int, path []int) {
	// fmt.Println("Currently at", curr, path)
	if curr == len(graph)-1 {
		*result = append(*result, path)
		return
	}
	if len(graph[curr]) == 0 {
		// fmt.Println("No more nodes at", curr)
		return
	}
	for _, node := range graph[curr] {
		newPath := make([]int, len(path))
		copy(newPath, path)
		dfs(node, graph, result, append(newPath, node))
	}
}