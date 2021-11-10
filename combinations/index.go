package algorithms

func Combine(n int, k int) [][]int {
	numbers := make([]int, n)
	result := make([][]int, 0)

	for i := 1; i <= n; i++ {
		numbers[i-1] = i
	}
	addCombination(&result, []int{}, 1, n, k)

	return result
}

func addCombination(result *[][]int, current []int, currentNum, n, k int) {
	if len(current) == k {
		newCurrent := make([]int, len(current))
		copy(newCurrent, current)
		*result = append(*result, newCurrent)
		return
	}
	for i := currentNum; i <= n; i++ {
		addCombination(result, append(current, i), i+1, n, k)
	}
}
