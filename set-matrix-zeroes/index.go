package algorithms

func setZeroes(matrix [][]int) [][]int {
	rows := make(map[int]bool)
	cols := make(map[int]bool)
	if len(matrix) == 0 {
		return matrix
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	for row := range rows {
		for i := range matrix[row] {
			matrix[row][i] = 0
		}
	}

	for col := range cols {
		for row := range matrix {
			matrix[row][col] = 0
		}
	}
	return matrix
}
