package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils"
)

func MaxArea(height []int) int {
	maxArea := 0
	l := 0
	r := len(height) - 1
	for l < r {
		area := (r - l) * Min(height[l], height[r])
		maxArea = Max(area, maxArea)
		if height[l] > height[r] {
			r -= 1
		} else {
			l += 1
		}
	}
	return maxArea
}
