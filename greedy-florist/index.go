package algorithms

import "sort"

func getMinimumCost(k int32, c []int32) int32 {
	count := 0
	total := 0
	cSorted := make([]int, len(c))
	for k, v := range c {
		cSorted[k] = int(v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(cSorted)))
	for i, v := range cSorted {
		if int32(i)%k == 0 {
			count++
		}
		total += v * count
	}
	return int32(total)
}
