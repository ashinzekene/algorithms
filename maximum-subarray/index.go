package algorithms

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prevMax := nums[0]
	maxVal := prevMax
	for _, v := range nums[1:] {
		prevMax = max(prevMax+v, v)
		maxVal = max(maxVal, prevMax)
	}
	return maxVal
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
