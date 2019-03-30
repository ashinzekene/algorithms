package algorithms

import "fmt"

func main() {
	fmt.Printf("Expected %d, got %d\n", 1, utopianTree(0))
	fmt.Printf("Expected %d, got %d\n", 2, utopianTree(1))
	fmt.Printf("Expected %d, got %d\n", 7, utopianTree(4))
	fmt.Printf("Expected %d, got %d\n", 14, utopianTree(5))
}

func utopianTree(n int32) int32 {
	var height int32 = 1
	for x := 1; int32(x) <= n; x++ {
		if x%2 == 1 {
			height *= 2
		} else {
			height++
		}
	}
	return height
}
