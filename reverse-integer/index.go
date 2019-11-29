package algorithms

import "math"

func reverse(x int) int {
	isNegative := false
	if x < 0 {
		isNegative = true
		x *= -1
	}
	reversed := 0
	for x > 0 {
		reversed = (reversed * 10) + (x % 10)
		x /= 10
	}
	if reversed >= math.MaxInt32 {
		return 0
	}
	if isNegative {
		return reversed * -1
	}
	return reversed
}
