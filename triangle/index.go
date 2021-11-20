package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils"
)

// https://leetcode.com/problems/triangle/

func MinimumTotal(triangle [][]int) int {
	dpTriangle := make([][]int, Min(2, len(triangle)))
	for i := range dpTriangle {
		dpTriangle[i] = make([]int, len(triangle[i]))
	}
	dpTriangle[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			current := triangle[i][j]
			if j == 0 {
				dpTriangle[1][j] = dpTriangle[0][j] + current
			} else if j == len(triangle[i])-1 {
				dpTriangle[1][j] = dpTriangle[0][j-1] + current
			} else {
				dpTriangle[1][j] = Min(
					dpTriangle[0][j-1],
					dpTriangle[0][j],
				) + current
			}
		}
		dpTriangle[0] = dpTriangle[1]
		dpTriangle[1] = make([]int, len(triangle[i])+1)
	}
	return Min(dpTriangle[0]...)
}
