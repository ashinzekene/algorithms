package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils"
)

func jump(nums []int) int {
	end := 0
	jumps := 0
	farthest := 0

	for i, v := range nums[:len(nums)-1] {
		farthest = Max(farthest, i+v)
		if i == end {
			end = farthest
			jumps++
		}
	}

	return jumps
}
