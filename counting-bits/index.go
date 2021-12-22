package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils"
)

func CountBits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	dp := make([]int, Max(n+1, 2))
	dp[0] = 0
	dp[1] = 1
	current := 1
	for i := 2; i <= n; i++ {
		if i == current*2 {
			current *= 2
		}
		nVal := i - current
		dp[i] = dp[nVal] + 1
	}
	return dp
}
