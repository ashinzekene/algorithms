package algorithms

import (
	. "github.com/ashinzekene/algorithms/largest-rectangle-in-histogram"
	. "github.com/ashinzekene/algorithms/utils"
)

func MaximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	result := 0
	rectangle := make([]int, len(matrix[0]))
	for _, row := range matrix {
		for i, val := range row {
			if val == '0' {
				rectangle[i] = 0
			} else {
				rectangle[i] += 1
			}
		}
		largestRectangle := LargestRectangleArea(rectangle)
		result = Max(result, largestRectangle)
	}
	return result
}
