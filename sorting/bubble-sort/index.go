package algorithms

func BubbleSort(arr []int) []int {
	isSorted := false
	unsortedLength := len(arr) - 1
	for !isSorted {
		isSorted = true
		for i := 0; i < unsortedLength; i++ {
			if arr[i] > arr[i+1] {
				isSorted = false
				temp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = temp
			}
		}
		unsortedLength--
	}
	return arr
}
