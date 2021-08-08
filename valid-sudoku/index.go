package algorithms

func isValidSudoku(board [][]byte) bool {
	// Validate horizontally
	for _, row := range board {
		byteIndex := make(map[byte]int)
		for _, b := range row {
			byteIndex[b]++
		}
		if !validateByteIndex(byteIndex) {
			return false
		}
	}

	// Validate vertically
	for i := range board[0] {
		byteIndex := make(map[byte]int)
		for _, col := range board {
			byteIndex[col[i]]++
		}
		if !validateByteIndex(byteIndex) {
			return false
		}
	}

	// Validate 3by3 matrix
	for count := 0; count < 9*3; count += 3 {
		jStart := (count / 9) * 3
		iStart := count % 9

		byteIndex := make(map[byte]int)
		for i := iStart; i < iStart+3; i++ {
			for j := jStart; j < jStart+3; j++ {
				byteIndex[board[i][j]]++
			}
		}
		if !validateByteIndex(byteIndex) {
			return false
		}
	}

	return true
}

func validateByteIndex(byteIndex map[byte]int) bool {
	for val, count := range byteIndex {
		if count > 1 && val != '.' {
			return false
		}
	}
	return true
}
