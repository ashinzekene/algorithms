package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils"
)

func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for diff := 1; diff <= n; diff++ {
		for currentN := 1; currentN <= n; currentN++ {
			if currentN < diff {
				continue
			} else if diff == 1 {
				dp[diff][currentN] = 0
			} else if diff == 2 {
				dp[diff][currentN] = currentN - 1
			} else if diff == 3 {
				dp[diff][currentN] = currentN - 1
			} else {
				start := currentN - diff + 1
				minValue := currentN * currentN // no guess can be larger than this
				mid := (start + currentN) / 2
				for guess := mid; guess <= currentN; guess++ {
					cost := guess + Max(dp[guess-start][guess-1], dp[currentN-guess][currentN])
					minValue = Min(minValue, cost)
				}
				dp[diff][currentN] = minValue
			}
		}
	}
	return dp[n][n]
}
