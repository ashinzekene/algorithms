package algorithms

func matchingStrings(strings []string, queries []string) []int32 {
	stringMap := make(map[string]int32)
	results := make([]int32, len(queries))
	for _, val := range strings {
		stringMap[val]++
	}
	for i, query := range queries {
		results[i] = stringMap[query]
	}
	return results
}
