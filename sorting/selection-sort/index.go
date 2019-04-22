package algorithms

func SelectionSort(A []int) []int {
	for x := 0; x < len(A); x++ {
		min := A[x]
		minIndex := x
		for i := x; i < len(A); i++ {
			if A[i] < min {
				min = A[i]
				minIndex = i
			}
		}
		temp := A[x]
		A[x] = A[minIndex]
		A[minIndex] = temp
	}
	return A
}
