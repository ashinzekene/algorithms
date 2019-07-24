package algorithms

import "math"

func Squares(a int32, b int32) int32 {
	var count int32
	var nextSquare int32
	val := a
	for {
		if val > b {
			return count
		}
		sqrt := int32(math.Sqrt(float64(val)))
		if sqrt*sqrt == val {
			count++
			nextSquare = (sqrt + 1) * (sqrt + 1)
			val = nextSquare
		} else {
			val++
		}
	}
}
