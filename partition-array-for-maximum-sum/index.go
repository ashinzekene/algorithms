package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils"
)

// algorithm
// e.g arr = [1,15,7,9,2,5,10], k = 3
// max(0, 1) = [1] == [1] = 1
// max(0, 2) = [1, 15] == [15, 15] = 30
// max(0, 3) = [1, 15, 7] == [15, 15, 15] = 45
// max(0, 4) = Max(
// 	max(0, 1) == 1 + max(1, 4) == [15,7,9] = [15, 15, 15] = 30 = 31
// 	max(0, 2) == 30 + max(2, 4) == [7,9] = [9, 9] = 18 = 48
// 	max(0, 3) == 45 + max(3, 4) == [9] = [9] = 9 = 54
// ) = 54
// etc

func MaxSumAfterPartitioning(arr []int, k int) int {
	maxes := make([]int, len(arr))
	maxes[0] = arr[0]

	for i := 1; i < k; i++ {
		maxes[i] = Max(arr[:i+1]...) * (i + 1)
	}
	for i := k; i < len(arr); i++ {
		max := 1
		for j := 0; j < k; j++ {
			current := maxes[i-k+j] + (k-j)*Max(arr[i-k+j+1:i+1]...)
			max = Max(
				max,
				current,
			)
		}
		maxes[i] = max
	}
	return maxes[len(maxes)-1]
}
