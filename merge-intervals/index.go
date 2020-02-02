package algorithms

import (
	"sort"

	. "github.com/ashinzekene/algorithms/utils"
)

func merge(intervals [][]int) [][]int {
	result := make([][]int, 0)
	if len(intervals) == 0 {
		return intervals
	}

	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		lastResult := result[len(result)-1]
		if lastResult[1] >= intervals[i][0] {
			lastResult[0] = Min(intervals[i][0], lastResult[0])
			lastResult[1] = Max(intervals[i][1], lastResult[1])
		} else {
			result = append(result, intervals[i])
		}
	}
	return result
}
