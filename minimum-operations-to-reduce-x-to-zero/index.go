package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils"
)

func minOperations(nums []int, x int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	expected := sum - x
	if expected < 0 {
		return -1
	}
	if expected == 0 {
		return len(nums)
	}
	result := len(nums) + 1
	cumm := 0
	start := 0
	for i := 0; i < len(nums); i++ {
		cumm += nums[i]
		for cumm > expected {
			cumm -= nums[start]
			start++
		}
		if cumm == expected {
			result = Min(result, (len(nums)-i-1)+start)
		}
	}
	if result == len(nums)+1 {
		return -1
	}
	return result
}
