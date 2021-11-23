package algorithms

func MergeSort(s []int) []int {
	l := len(s)
	if l <= 1 {
		return s
	} else if l == 2 {
		if s[0] > s[1] {
			return []int{s[1], s[0]}
		}
		return s
	}
	m := l / 2
	a, b := s[:m], s[m:]
	sorted_a := MergeSort(a)
	sorted_b := MergeSort(b)
	return mergeSortedList(sorted_a, sorted_b)
}

func mergeSortedList(a, b []int) []int {
	if len(a) == 0 {
		return b
	} else if len(b) == 0 {
		return a
	}
	if a[0] < b[0] {
		return append([]int{a[0]}, mergeSortedList(a[1:], b)...)
	} else {
		return append([]int{b[0]}, mergeSortedList(b[1:], a)...)
	}
}
