package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils"
)

func GetMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 0
	}
	arr := make([]int, n+1)
	arr[0] = 0
	arr[1] = 1
	max := arr[1]
	for i := 2; i <= n; i++ {
		div := i / 2
		result := arr[div]
		if i%2 == 1 {
			result += arr[div+1]
		}
		arr[i] = result
		max = Max(result, max)
	}
	return max
}
