package algorithms

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Expected %d got %d \n", 0, viralAvertising(0))
	fmt.Printf("Expected %d got %d \n", 2, viralAvertising(1))
	fmt.Printf("Expected %d got %d \n", 5, viralAvertising(2))
	fmt.Printf("Expected %d got %d \n", 24, viralAvertising(5))
}

func viralAvertising(n int32) int32 {
	shared := 5
	liked := 0

	for n > 0 {
		likes := int(math.Floor(float64(shared / 2)))
		shared = likes * 3
		liked += likes
		n--
	}
	return int32(liked)
}
