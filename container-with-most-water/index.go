package algorithms

import "math"

func maxArea(height []int) int {
	maxArea := 0
	l := 0
	r := len(height) - 1
	for l < r {
		area := float64(r-l) * math.Min(float64(height[l]), float64(height[r]))
		maxArea = int(math.Max(float64(area), float64(maxArea)))
		if height[l] > height[r] {
			r -= 1
		} else {
			l += 1
		}
	}
	return maxArea
}
