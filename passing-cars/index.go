package algorithms

func PassingCars(A []int) int {
	westMovements := make([]int, len(A))
	l := len(A)
	result := 0

	westMovements[l-1] = A[l-1]
	for i := l - 2; i >= 0; i-- {
		westMovements[i] = A[i] + westMovements[i+1]
	}
	for i, m := range A {
		if m == 0 {
			result += westMovements[i]
		}
		if result > 1000000000 {
			return -1
		}
	}
	return result
}
