package algorithms

func countTriplets(arr []int64, r int64) int64 {
	integerMap := make(map[int64]int64, 0)
	var result int64 = 0
	for _, v := range arr {
		integerMap[v]++
	}
	for k, v := range integerMap {
		if v > 0 {
			result += integerMap[k] * integerMap[k*r] * integerMap[k*r*r]
		}
	}
	return result
}
