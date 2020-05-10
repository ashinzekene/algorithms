package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils"
)

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dpMatrix := make([][]int, m)
	for i := 0; i < m; i++ {
		dpMatrix[i] = make([]int, n)
	}
	dpMatrix[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		dpMatrix[0][i] = grid[0][i] + dpMatrix[0][i-1]
	}
	for i := 1; i < m; i++ {
		dpMatrix[i][0] = grid[i][0] + dpMatrix[i-1][0]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			prevVal := Min(dpMatrix[i-1][j], dpMatrix[i][j-1])
			dpMatrix[i][j] = prevVal + grid[i][j]
		}
	}
	return dpMatrix[m-1][n-1]
}
