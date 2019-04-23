package algorithms

import (
	"math"
)

func jumpingOnClouds(c []int32, k int32) int32 {
	jumps := int(math.Ceil(float64(len(c)) / float64(k)))
	var energy int32 = 100
	for x := 1; x <= jumps; x++ {
		energy--
		i := x * int(k)
		if c[i%len(c)] == 1 {
			energy -= 2
		}
	}
	return energy
}
