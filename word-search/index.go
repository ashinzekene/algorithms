package algorithms

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				prevVal := board[i][j]
				board[i][j] = '-'
				if findWord(board, word[1:], i, j) {
					return true
				}
				board[i][j] = prevVal
			}
		}
	}
	return false
}

func findWord(board [][]byte, wordLeft string, pI, pJ int) bool {
	if wordLeft == "" {
		return true
	}
	if pI-1 >= 0 && board[pI-1][pJ] == wordLeft[0] {
		prevVal := board[pI-1][pJ]
		board[pI-1][pJ] = '-'
		if findWord(board, wordLeft[1:], pI-1, pJ) {
			return true
		}
		board[pI-1][pJ] = prevVal

	}
	if pJ-1 >= 0 && board[pI][pJ-1] == wordLeft[0] {
		prevVal := board[pI][pJ-1]
		board[pI][pJ-1] = '-'
		if findWord(board, wordLeft[1:], pI, pJ-1) {
			return true
		}
		board[pI][pJ-1] = prevVal

	}
	if pJ+1 < len(board[0]) && board[pI][pJ+1] == wordLeft[0] {
		prevVal := board[pI][pJ+1]
		board[pI][pJ+1] = '-'
		if findWord(board, wordLeft[1:], pI, pJ+1) {
			return true
		}
		board[pI][pJ+1] = prevVal

	}
	if pI+1 < len(board) && board[pI+1][pJ] == wordLeft[0] {
		prevVal := board[pI+1][pJ]
		board[pI+1][pJ] = '-'
		if findWord(board, wordLeft[1:], pI+1, pJ) {
			return true
		}
		board[pI+1][pJ] = prevVal
	}
	return false
}
