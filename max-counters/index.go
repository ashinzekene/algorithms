package algorithms

func Solution(N int, A []int) []int {
	operations := make(map[int]int, N+1)
	result := make([]int, N)
	max := 0
	generalMax := 0
	lastMaxIndex := 0
	// Get last max index
	for i := len(A) - 1; i >= 0; i-- {
		if A[i] > N {
			lastMaxIndex = i
			break
		}
	}

	// Get the last general max
	for _, op := range A {
		if op > N {
			operations[N+1] = max
			generalMax = max
			continue
		}
		if operations[op] <= generalMax {
			operations[op] = generalMax + 1
		} else {
			operations[op]++
		}
		if operations[op] > max {
			max = operations[op]
		}
	}

	// use the last operations after
	operations = make(map[int]int)
	for i := lastMaxIndex; i < len(A); i++ {
		operations[A[i]]++
	}

	for i := 1; i <= N; i++ {
		result[i-1] = operations[i] + generalMax
	}

	return result
}
