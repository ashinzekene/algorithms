package algorithms

import . "github.com/ashinzekene/algorithms/utils"

func canJump(nums []int) bool {
	max := 0
	for i := 0; i < len(nums)-1; i++ {
		max = Max(max-1, nums[i])
		if max == 0 {
			return false
		}
	}
	return true
}
