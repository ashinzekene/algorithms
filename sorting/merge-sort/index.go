package algorithms

import "fmt"

func main() {
	fmt.Println(merge_sort([]int{1,4,5,2,6,7,9,8,3,24,72,14,13,55,42,16,31,30,29,28,27,20, 10, 12, 11}))
}

func merge_arr(L []int, R []int, A []int) []int {
	l, r, i := 0, 0, 0
	for ;l < len(L) && r < len(R); {
		if L[l] < R[r] {
			A[i] = L[l]
			l++
		} else {
			A[i] = R[r]
			r++
		}
		i++
	}
	for ;l <len(L); {
		A[i] = L[l]
		l++
		i++
	}
	for ;r <len(R); {
		A[i] = R[r]
		r++
		i++
	}
	return A
}

func merge_sort(A []int) []int {
	if len(A) < 2 {
		return A
	}
	mid := len(A) / 2
	// append([]int{}, A[:mid]...) - creates a new slice without referencing A
	L := merge_sort(append([]int{}, A[:mid]...))
	R := merge_sort(append([]int{}, A[mid:]...))
	return merge_arr(L, R, A)
}
