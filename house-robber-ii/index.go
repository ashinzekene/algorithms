package algorithms

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return Max(rob2(nums[1:]), rob2(nums[:len(nums)-1]))
}

func rob2(nums []int) int {
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

func Max(a ...int) int {
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}
