package algorithms

import . "github.com/ashinzekene/algorithms/utils"

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prevMax := nums[0]
	prevMin := nums[0]
	maxProduct := nums[0]
	minProduct := nums[0]
	for _, num := range nums[1:] {
		newMaxProduct := Max(maxProduct*num, minProduct*num, num)
		minProduct = Min(minProduct*num, maxProduct*num, num)
		maxProduct = newMaxProduct

		prevMax = Max(prevMax, maxProduct)
		prevMin = Min(prevMin, minProduct)
	}
	return prevMax
}
