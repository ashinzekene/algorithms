package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils"
)

func LargestRectangleArea(heights []int) int {
	maxArea := 0
	stack := make([][2]int, 0)
	for i, height := range heights {
		currentIndex := i
		if len(stack) > 0 && height < stack[len(stack)-1][0] {
			for len(stack) > 0 && height < stack[len(stack)-1][0] {
				v := stack[len(stack)-1][0]
				index := stack[len(stack)-1][1]
				length := i - index
				stack = stack[:len(stack)-1]
				area := v * length
				maxArea = Max(maxArea, area)
				currentIndex = index
			}
		}
		stack = append(stack, [2]int{height, currentIndex})
	}
	l := len(heights)
	for len(stack) > 0 {
		current := stack[len(stack)-1]
		area := current[0] * (l - current[1])
		maxArea = Max(maxArea, area)
		stack = stack[:len(stack)-1]
	}
	return maxArea
}
