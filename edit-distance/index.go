package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils"
)

func MinDistance(word1 string, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)
	dpMatrix := make([][]int, l1+1)
	for i := range dpMatrix {
		dpMatrix[i] = make([]int, l2+1)
	}
	dpMatrix[0][0] = 0
	for i := 1; i <= l1; i++ {
		dpMatrix[i][0] = dpMatrix[i-1][0] + 1
	}
	for j := 1; j <= l2; j++ {
		dpMatrix[0][j] = dpMatrix[0][j-1] + 1
	}
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			w1 := word1[i-1]
			w2 := word2[j-1]
			if w1 == w2 {
				dpMatrix[i][j] = dpMatrix[i-1][j-1]
			} else {
				dpMatrix[i][j] = Min(
					dpMatrix[i-1][j],
					dpMatrix[i][j-1],
					dpMatrix[i-1][j-1],
				) + 1
			}
		}
	}
	return dpMatrix[l1][l2]
}
