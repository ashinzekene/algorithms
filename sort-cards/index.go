package algorithms

import "fmt"

func main () {
	Solution("A586QT", "JJ653K")
	Solution("23A84Q", "K2Q25J")
}

func Solution(A string, B string) int {
	var count int = 0
	ratings := map[string]int{
		"A": 0,
		"K": 1,
		"Q": 2,
		"J": 3,
		"T": 4,
	}

	for x:= 0; x < len(A); x++ {
		if A[x] >64 && B[x] > 64 {
			if ratings[string(A[x])] < ratings[string(B[x])] {
				count++
			}
		} else if A[x] > B[x] {
			count++
		}
	}
	fmt.Println(count)
	return count
}