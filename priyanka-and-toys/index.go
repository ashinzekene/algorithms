package algorithms

import (
	"sort"
)

func toys(w []int32) int32 {
	sort.Slice(w, func(i, j int) bool {
		return w[i] < w[j]
	})
	mins := make([]int32, 0)
	for i, v := range w {
		if i == 0 {
			mins = append([]int32{v}, mins...)
		} else {
			if v > mins[0]+4 {
				mins = append([]int32{v}, mins...)
			}
		}
	}
	return int32(len(mins))
}

func toys1(nums []int32) int32 {
	if len(nums) == 0 {
		return 0
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var containers int32 = 1
	minWeight := nums[0]
	for _, v := range nums[1:] {
		if v-minWeight > 4 {
			minWeight = v
			containers++
		}
	}
	return containers
}
