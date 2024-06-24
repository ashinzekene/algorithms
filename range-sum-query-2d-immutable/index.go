package algorithms

type NumMatrix struct {
	matrix [][]int
}


func Constructor(matrix [][]int) NumMatrix {
	newMatrix := make([][]int, len(matrix)+1)
	for i := range newMatrix {
		newMatrix[i] = make([]int, len(matrix[0])+1)
		for j := 1; j< len(newMatrix[i]); j++ {
			if i == 0 {
				continue
			}
			newMatrix[i][j] = newMatrix[i-1][j] + newMatrix[i][j-1] + matrix[i-1][j-1] - newMatrix[i-1][j-1]
		}
	}
	return NumMatrix{
		matrix: newMatrix,
	}
}

func (nm *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	total := nm.matrix[row2+1][col2+1]
	left := nm.matrix[row1][col2+1]
	top := nm.matrix[row2+1][col1]
	head := nm.matrix[row1][col1]
	return total - left - top + head
}
