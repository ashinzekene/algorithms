package algorithms

func numIslands(grid [][]byte) int {
	islandCount := 0
	for i := range grid {
		for j := range grid[0] {
			if island(grid, i, j) {
				islandCount++
			}
		}
	}
	return islandCount
}

func island(grid [][]byte, i, j int) bool {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '#' || grid[i][j] == '0' {
		return false
	}
	grid[i][j] = '#'
	island(grid, i-1, j)
	island(grid, i, j-1)
	island(grid, i+1, j)
	island(grid, i, j+1)
	return true
}
