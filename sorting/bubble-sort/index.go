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

// Faster bubble sort
func bubbleSort1(arr []int) []int {
	end := 1
	for x := 0; x < len(arr); x++ {
		for i := 0; i < len(arr)-end; i++ {
			if arr[i] > arr[i+1] {
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
			}
			end++
		}
	}
	return arr
}
