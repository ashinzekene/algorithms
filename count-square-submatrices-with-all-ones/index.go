package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils"
)

func CountSquares(matrix [][]int) int {
	result := 0
	oldDp := make([]int, len(matrix[0]))
	newDp := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			current := matrix[i][j]
			if i == 0 || j == 0 {
				newDp[j] = current
			} else {
				val := (Min(newDp[j], newDp[j-1], oldDp[j-1]) + current) * current
				newDp[j] = val
			}
			result += newDp[j]
		}
		copy(oldDp, newDp)
	}
	return result
}
