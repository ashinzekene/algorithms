package algorithms

func isValidSudoku(board [][]byte) bool {
	for i, row := range board {
		// Check if colum ns are valid
		validCol := isValidCol(row)
		if !validCol {
			return false
		}
		validRow := checkValidRow(board, i)
		if !validRow {
			return false
		}
	}
	var cell [3][3]byte
	for l := 0; l < len(board); l += 3 {
		for k := 0; k < len(board); k += 3 {
			for i := l; i < l+3; i++ {
				for j := k; j < k+3; j++ {
					cell[i%3][j%3] = board[i][j]
				}
			}
			valid := checkValid3Cell(cell)
			if !valid {
				return false
			}
		}
	}
	return true
}

func isValidCol(col []byte) bool {
	byteMap := make(map[byte]bool)
	for _, num := range col {
		if byteMap[num] && num != '.' {
			return false
		}
		byteMap[num] = true
	}
	return true
}

func checkValidRow(board [][]byte, i int) bool {
	byteMap := make(map[byte]bool)
	for _, col := range board {
		if byteMap[col[i]] && col[i] != '.' {
			return false
		}
		byteMap[col[i]] = true
	}
	return true
}

func checkValid3Cell(cell [3][3]byte) bool {
	byteMap := make(map[byte]bool)
	for i := 0; i < len(cell); i++ {
		for j := 0; j < len(cell[0]); j++ {
			val := cell[i][j]
			if byteMap[val] && val != '.' {
				return false
			}
			byteMap[val] = true
		}
	}
	return true
}
