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

// Interesting nlog(n) solution
func LargestRectangleArea2(heights []int) int {
	if len(heights) == 0 {
		return 0
	} else if len(heights) == 1 {
		return heights[0]
	}
	min, minIndex := getMin(heights)
	return Max(
		LargestRectangleArea2(heights[:minIndex]),
		len(heights)*min,
		LargestRectangleArea2(heights[minIndex+1:]),
	)
}

func getMin(heights []int) (int, int) {
	min := heights[0]
	minIndex := 0
	for i, height := range heights {
		if height < min {
			min = height
			minIndex = i
		}
	}
	return min, minIndex
}
