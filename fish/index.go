package algorithms

func Fish(A []int, B []int) int {
	stack := make([][2]int, 0)
	for i := 0; i < len(A); {
		weight := A[i]
		dir := B[i]
		if len(stack) == 0 {
			stack = append(stack, [2]int{weight, B[i]})
		} else {
			lastWeight := stack[len(stack)-1][0]
			lastDir := stack[len(stack)-1][1]
			if lastDir == dir || lastDir == 0 {
				stack = append(stack, [2]int{weight, B[i]})
			} else if lastWeight < weight {
				stack = stack[:len(stack)-1]
				i--
			}
		}
		i++
	}
	return len(stack)
}
