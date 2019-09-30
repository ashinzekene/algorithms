package algorithms

import (
	"math"
	"sort"
)

func minimumAbsoluteDifference(arr []int32) int32 {
	arrInts := make([]int, len(arr))
	for i, v := range arr {
		arrInts[i] = int(v)
	}
	sort.Ints(arrInts)
	minDiff := arrInts[len(arrInts)-1] - arrInts[0]
	for i, v := range arrInts {
		if i == len(arrInts)-1 {
			break
		}
		diff := int(math.Abs(float64(v - arrInts[i+1])))
		if diff < minDiff {
			minDiff = diff
		}
	}
	return int32(minDiff)
}
