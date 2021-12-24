package algorithms

func solve(board [][]byte) {
	n := len(board)
	m := len(board[0])
	// top to bottom
	for i := 0; i < n; i++ {
		checkZeros(board, i, 0)
		checkZeros(board, i, m-1)
	}
	// left to right
	for j := 0; j < m; j++ {
		checkZeros(board, 0, j)
		checkZeros(board, n-1, j)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			} else if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}
}

func checkZeros(board [][]byte, i, j int) {
	n := len(board)
	m := len(board[0])
	if i < 0 || i > n-1 || j < 0 || j > m-1 {
		return
	}
	if board[i][j] != 'O' {
		return
	}
	board[i][j] = '#'
	checkZeros(board, i-1, j)
	checkZeros(board, i+1, j)
	checkZeros(board, i, j-1)
	checkZeros(board, i, j+1)
}
