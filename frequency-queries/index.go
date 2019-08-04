package algorithms

func FreqQuery(queries [][]int32) []int32 {
	intMap := make(map[int32]int32)
	countMap := make(map[int32]int32)
	result := make([]int32, 0, len(queries))
	for _, val := range queries {
		if val[0] == 1 {
			if intMap[val[1]] > 0 {
				countMap[intMap[val[1]]]--
			}
			intMap[val[1]]++
			countMap[intMap[val[1]]]++
		}
		if val[0] == 2 {
			if intMap[val[1]] > 0 {
				countMap[intMap[val[1]]]--
				intMap[val[1]]--
				countMap[intMap[val[1]]]++
			}
		}
		if val[0] == 3 {
			count := 0
			if countMap[val[1]] > 0 {
				count++
			}
			result = append(result, int32(count))
		}
	}
	return result
}
