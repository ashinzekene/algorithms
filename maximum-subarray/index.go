package algorithms

import "github.com/ashinzekene/algorithms/utils"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prevMax := nums[0]
	maxVal := prevMax
	for _, v := range nums[1:] {
		prevMax = utils.Max(prevMax+v, v)
		maxVal = utils.Max(maxVal, prevMax)
	}
	return maxVal
}
