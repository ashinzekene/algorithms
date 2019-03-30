package algorithms

import (
	"fmt"
)

func main() {
	fmt.Println(pickingNumbers([]int32{1, 1, 2, 2, 4, 4, 5, 5, 5}))
	fmt.Println(pickingNumbers([]int32{4, 6, 5, 3, 3, 1}))
	fmt.Println(pickingNumbers([]int32{1, 2, 2, 3, 1, 2}))
}

func pickingNumbers(a []int32) int32 {
	var maxCount int32 = 0
	var lowerCount int32 = 0
	var upperCount int32 = 0

	for x := 0; x < len(a); x++ {
		for i := 0; i < len(a); i++ {
			if a[x]-a[i] == 1 {
				lowerCount++
			} else if a[i]-a[x] == 1 {
				upperCount++
			} else if a[x] == a[i] {
				upperCount++
				lowerCount++
			}
		}
		if upperCount > maxCount {
			maxCount = upperCount
		}
		if lowerCount > maxCount {
			maxCount = lowerCount
		}
		upperCount = 0
		lowerCount = 0
	}
	return maxCount
}
