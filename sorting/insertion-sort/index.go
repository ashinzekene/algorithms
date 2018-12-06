package main

import "fmt"

func main() {
	fmt.Println(insertionSort([]int{1,9,8,4,6,7,3,12,6,3,2,22}))
}

func insertionSort(arr []int) []int {
	for x := 1; x < len(arr); x++ {
		item := arr[x]
		i := x - 1
		for ;i >= 0 && item < arr[i]; i-- {
			arr[i + 1] = arr[i]
		}
		arr[i + 1] = item
	}
	return arr
}