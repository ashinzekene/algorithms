package algorithms

import . "github.com/ashinzekene/algorithms/utils"

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	result := 0
	dpMatrix := make([][]int, len(matrix))
	for i := range dpMatrix {
		dpMatrix[i] = make([]int, len(matrix[0]))
	}
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := len(matrix[0]) - 1; j >= 0; j-- {
			dpMatrix[i][j] = getCount(dpMatrix, matrix[i][j], i, j)
			result = Max(int(dpMatrix[i][j]), result)
		}
	}
	return result * result
}

func getCount(dpMatrix [][]int, v byte, i, j int) int {
	var right, down, side int
	if j+1 < len(dpMatrix[0]) {
		right = int(dpMatrix[i][j+1])
	}
	if i+1 < len(dpMatrix) {
		down = int(dpMatrix[i+1][j])
	}
	if j+1 < len(dpMatrix[0]) && i+1 < len(dpMatrix) {
		side = int(dpMatrix[i+1][j+1])
	}
	return (int(v) - 48) * (Min(right, down, side) + 1)
}
