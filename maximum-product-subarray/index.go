package algorithms

import "github.com/ashinzekene/algorithms/utils"

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	mainMax := nums[0]
	prevMin := mainMax
	prevMax := mainMax
	for _, v := range nums[1:] {
		max_min := utils.Max(prevMin*v, v)
		min_min := utils.Min(prevMin*v, v)
		max_max := utils.Max(prevMax*v, v)
		min_max := utils.Min(prevMax*v, v)

		prevMin = utils.Min(min_max, min_min)
		prevMax = utils.Max(max_min, max_max)

		mainMax = utils.Max(mainMax, prevMax)
	}
	return mainMax
}
