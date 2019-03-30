package algorithms

import "fmt"

func main() {
	fmt.Println(selection_sort([]int{1,4,5,2,6,7,9,8,3,24,72,14,13,55,42,31,30,29,28,27,20}))
}

func selection_sort(A []int) []int {
	for x := 0; x < len(A); x++ {
		min := A[x];
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