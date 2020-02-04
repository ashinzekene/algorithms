package algorithms

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	usedCells := make(map[int]map[int]bool)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				hasWord := findNextChar(board, i, j, word[1:], usedCells)
				if hasWord {
					return true
				}
			}
		}
	}
	return false
}

func findNextChar(grid [][]byte, i, j int, word string, usedCells map[int]map[int]bool) bool {
	if usedCells[i] == nil {
		usedCells[i] = make(map[int]bool)
	}
	usedCells[i][j] = true
	if word == "" {
		return true
	}
	if i-1 >= 0 && !usedCells[i-1][j] {
		if grid[i-1][j] == word[0] && findNextChar(grid, i-1, j, word[1:], usedCells) {
			return true
		}
	}
	if j+1 < len(grid[0]) && !usedCells[i][j+1] {
		if grid[i][j+1] == word[0] && findNextChar(grid, i, j+1, word[1:], usedCells) {
			return true
		}
	}
	if i+1 < len(grid) && !usedCells[i+1][j] {
		if grid[i+1][j] == word[0] && findNextChar(grid, i+1, j, word[1:], usedCells) {
			return true
		}
	}
	if j-1 >= 0 && !usedCells[i][j-1] {
		if grid[i][j-1] == word[0] && findNextChar(grid, i, j-1, word[1:], usedCells) {
			return true
		}
	}
	usedCells[i][j] = false
	return false
}
