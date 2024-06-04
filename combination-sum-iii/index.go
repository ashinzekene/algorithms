package algorithms

func combinationSum3(k int, n int) [][]int {
	result := make([][]int, 0)
	for i := 1; i < 10;i++ {
			getSum([]int{i}, k, n-i, &result)
	}
	return result
}

func getSum(used []int, k, n int, result *[][]int) {
	if n == 0 && len(used) == k {
		*result = append(*result, append([]int{}, used...))
		return
	}

	if n < 0 || len(used) >= k {
			return
	}

	for i := used[len(used)-1]+1; i < 10;i++ {
			if !isIn(used, i) {
					getSum(append(used, i), k, n-i, result)
			}
	}
}

func isIn(arr []int, val int) bool {
	for _, curr := range arr {
			if curr == val {
					return true
			}
	}
	return false
}
