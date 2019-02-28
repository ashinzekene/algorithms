package main

import "fmt"

func main() {
	fmt.Println("Result", birthday([]int32{1, 2, 1, 3, 2}, 3, 2))
	fmt.Println("Result", birthday([]int32{1, 1, 1, 1, 1}, 3, 2))
	fmt.Println("Result", birthday([]int32{4}, 4, 1))
	fmt.Println("Result", birthday([]int32{2, 5, 1, 3, 4, 4, 3, 5, 1, 1, 2, 1, 4, 1, 3, 3, 4, 2, 1}, 18, 7))
	fmt.Println("Result", birthday([]int32{3, 5, 4, 1, 2, 5, 3, 4, 3, 2, 1, 1, 2, 4, 2, 3, 4, 5, 3, 1, 2, 5, 4, 5, 4, 1, 1, 5, 3, 1, 4, 5, 2, 3, 2, 5, 2, 5, 2, 2, 1, 5, 3, 2, 5, 1, 2, 4, 3, 1, 5, 1, 3, 3, 5}, 18, 6))
}

func birthday(s []int32, d int32, m int32) int32 {
	var result int32 = 0
	var array = make([]int32, 0, 3)
	var sum int32 = 0

	for i := 0; i < len(s); i++ {
		array = []int32{s[i]}
		sum = s[i]
		if sum == d && len(array) == int(m) {
			result++
		}
		for j := i + 1; j < len(s); j++ {
			sum += s[j]
			array = append(array, s[j])
			if sum == d && len(array) == int(m) {
				result++
			}
			if sum > d || len(array) > int(m) {
				// Exit inner loop
			}
		}

	}
	return result
}
