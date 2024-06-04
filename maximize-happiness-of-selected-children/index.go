package algorithms

import "github.com/ashinzekene/algorithms/utils/heaps"

func maximumHappinessSum(happiness []int, k int) int64 {
	h := heaps.NewMaxHeap(len(happiness))
	for _, i := range happiness {
		h.Insert(i)
	}
	var result int64 = 0
	for i := 0; i < k; i++ {
		result += int64(max(h.RemoveMax() - i, 0))
	}
	return result
}