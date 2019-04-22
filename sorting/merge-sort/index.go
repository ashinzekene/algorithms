package algorithms

func merge_arr(L []int, R []int, A []int) []int {
	l, r, i := 0, 0, 0
	for l < len(L) && r < len(R) {
		if L[l] < R[r] {
			A[i] = L[l]
			l++
		} else {
			A[i] = R[r]
			r++
		}
		i++
	}
	for l < len(L) {
		A[i] = L[l]
		l++
		i++
	}
	for r < len(R) {
		A[i] = R[r]
		r++
		i++
	}
	return A
}

func MergeSort(A []int) []int {
	if len(A) < 2 {
		return A
	}
	mid := len(A) / 2
	// append([]int{}, A[:mid]...) - creates a new slice without referencing A
	L := MergeSort(append([]int{}, A[:mid]...))
	R := MergeSort(append([]int{}, A[mid:]...))
	return merge_arr(L, R, A)
}
