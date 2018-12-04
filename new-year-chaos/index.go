package main

import "fmt"

// 12345678
// 13247856
// 13247568

func main() {
	newYearChaos([]int{1,2,3,4,5,6,7});
}

func newYearChaos(A []int) {
	Y := append(A, 8)
	for _, x := range Y {
		fmt.Println(x)
	}
}