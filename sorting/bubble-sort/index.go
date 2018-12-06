package main

import "fmt"

func main() {
	fmt.Println(bubbleSort([]int{1,4,5,2,6,7,9,8,3,24,72,14,13,55,42,31,30,29,28,27,20}))
}

func bubbleSort(arr []int) []int {
	for x := 0; x < len(arr); x++ {
    for i := 0; i < len(arr) - 1;i++ {
      if arr[i] > arr[i + 1] {
        temp := arr[i]
        arr[i] = arr[i + 1]
        arr[i + 1] = temp
      }
    }
  }
  return arr
}