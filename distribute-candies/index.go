package algorithms

func distributeCandies(candies []int) int {
	maxSisCandies := len(candies) / 2
	kindMap := make(map[int]bool)
	kinds := 0
	for _, v := range candies {
		if kinds == maxSisCandies {
			return kinds
		}
		if !kindMap[v] {
			kinds++
			kindMap[v] = true
		}
	}
	return kinds
}
