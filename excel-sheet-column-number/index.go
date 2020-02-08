package algorithms

import "math"

func titleToNumber(s string) int {
	result := 0
	for i := len(s) - 2; i >= 0; i-- {
		alpha := int(s[i]) - 64
		multiplier := math.Pow(26, float64(len(s)-i-1))
		result += int(multiplier) * alpha
	}
	result += int(s[len(s)-1]) - 64
	return result
}
