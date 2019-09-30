package algorithms

import "sort"

func luckBalance(k int32, contests [][]int32) int32 {
	important := make([]int, 0)
	var total int32 = 0
	for _, v := range contests {
		total += v[0]
		if v[1] == 1 {
			important = append(important, int(v[0]))
		}
	}
	sort.Ints(important)
	wins := len(important) - int(k)
	for i, v := range important {
		if i+1 > wins {
			return total
		}
		total -= int32(v * 2)
	}
	return total
}
