package algorithms

import . "github.com/ashinzekene/algorithms/utils"

func TapeEquilibrium(A []int) int {
	totalSum := 0
	for _, num := range A {
		totalSum += num
	}
	sum := 0
	diff := Abs((totalSum - A[0]) - A[0])
	shouldContinue := true
	for i := 0; shouldContinue && i < len(A)-1; i++ {
		sum += A[i]
		newDiff := Abs((totalSum - sum) - sum)
		diff = Min(newDiff, diff)
	}
	return diff
}
