package algorithms

import "sort"

func CutTheSticks(arr []int32) []int32 {
	res := make(map[int]int32)
	result := make([]int32, 0, len(arr))
	l := int32(len(arr))
	for _, i := range arr {
		res[int(i)]++
	}
	keys := make([]int, 0, len(res))
	for key := range res {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, i := range keys {
		result = append(result, l)
		l = l - res[i]
	}
	return result
}
