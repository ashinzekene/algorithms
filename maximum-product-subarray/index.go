package algorithms

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	mainMax := nums[0]
	prevMin := mainMax
	prevMax := mainMax
	for _, v := range nums[1:] {
		max_min := max(prevMin*v, v)
		min_min := min(prevMin*v, v)
		max_max := max(prevMax*v, v)
		min_max := min(prevMax*v, v)

		prevMin = min(min_max, min_min)
		prevMax = max(max_min, max_max)

		mainMax = max(mainMax, prevMax)
	}
	return mainMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
