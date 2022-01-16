package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils"
)

func rob(nums []int) int {
	dp := make([]int, len(nums))
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp[0] = nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		current := nums[i]
		if i > 1 {
			current += dp[i-2]
		}
		max = Max(current, max)
		dp[i] = max
	}
	return dp[len(nums)-1]
}

// O(1) space
func rob1(nums []int) int {
	prevCumm := 0
	cumm := 0
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		current := nums[i]
		newCumm := current + prevCumm
		max = Max(newCumm, cumm)
		prevCumm = Max(cumm, prevCumm)
		cumm = newCumm
	}
	return max
}
