package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils"
)

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i := 3; i <= n; i++ {
		mid := i / 2
		for j := 1; j <= mid; j++ {
			a := Max(j, dp[j])
			b := Max(i-j, dp[i-j])
			dp[i] = Max(dp[i], a*b)
		}
	}
	return dp[n]
}
