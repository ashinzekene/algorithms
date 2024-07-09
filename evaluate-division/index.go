package algorithms

import "math"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
  adjList := map[string]map[string]float64{}
	results := make([]float64, len(queries))
	for i, equation := range equations {
		valA := equation[0]
		valB := equation[1]
	
		if adjList[valA] == nil {
			adjList[valA] = map[string]float64{}
		}
		adjList[valA][valB] = values[i]

		if adjList[valB] == nil {
			adjList[valB] = map[string]float64{}
		}
		if values[i] == 0 {
			adjList[valB][valA] = math.Inf(-1)
		} else {
			adjList[valB][valA] = 1/values[i]
		}

	}

	for i, query := range queries {
		start := query[0]
		end := query[1]
		visited := map[string]bool{}
		ans := dfs(start, end, visited, adjList)
		results[i] = ans
	}

	return results
}

func dfs(current, end string, visited map[string]bool, adjList map[string]map[string]float64) float64 {
	if visited[current] || adjList[current] == nil {
		return -1
	}
	if current == end {
		return 1.0
	}
	visited[current] = true
	for child, value := range adjList[current] {
		current := value * dfs(child, end, visited, adjList)
		if current >= 0 {
			return current
		}
	}
	visited[current] = false
	return -1
}

