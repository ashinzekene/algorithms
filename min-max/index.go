package algorithms

import "sort"

func maxMin(k int32, arr []int32) int32 {
	kVal := int(k - 1)
	sortedArr := make([]int, len(arr))
	for i, v := range arr {
		sortedArr[i] = int(v)
	}
	sort.Ints(sortedArr)
	minDiff := sortedArr[kVal] - sortedArr[0]
	for i := 0; i < len(sortedArr)-kVal; i++ {
		diff := sortedArr[i+kVal] - sortedArr[i]
		if diff < minDiff {

			minDiff = diff
		}
	}
	return int32(minDiff)
}
