package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils"
)

func minSubArrayLen(target int, nums []int) int {
	i := 0
	j := 0
	minLength := len(nums) + 1
	currentSum := nums[0]
	for i < len(nums) && j < len(nums) {
		if !(currentSum < target) {
			minLength = Min(j-i+1, minLength)
			currentSum -= nums[i]
			i++
		} else if !(currentSum > target) {
			j++
			if j < len(nums) {
				currentSum += nums[j]
			}
		}
	}
	if minLength > len(nums) {
		return 0
	}
	return minLength
}
