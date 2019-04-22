package algorithms

func BubbleSort(arr []int) []int {
	for x := 0; x < len(arr); x++ {
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
			}
		}
	}
	return arr
}