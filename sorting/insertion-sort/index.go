package algorithms

func InsertionSort(arr []int) []int {
	for x := 1; x < len(arr); x++ {
		item := arr[x]
		i := x - 1
		for ; i >= 0 && item < arr[i]; i-- {
			arr[i+1] = arr[i]
		}
		arr[i+1] = item
	}
	return arr
}
