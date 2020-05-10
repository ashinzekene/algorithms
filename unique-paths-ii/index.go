package algorithms

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dpMatrix := make([][]int, m)
	for i := 0; i < m; i++ {
		dpMatrix[i] = make([]int, n)
	}
	// prefill top row
	zeroPresent := false
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 0 && !zeroPresent {
			dpMatrix[0][i] = 1
		} else {
			zeroPresent = true
		}
	}
	// Prefill first column
	zeroPresent = false
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 0 && !zeroPresent {
			dpMatrix[i][0] = 1
		} else {
			zeroPresent = true
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dpMatrix[i][j] = 0
			} else {
				dpMatrix[i][j] = dpMatrix[i][j-1] + dpMatrix[i-1][j]
			}
		}
	}
	return dpMatrix[m-1][n-1]
}
