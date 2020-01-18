package algorithms

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func trap(height []int) int {
	hLen := len(height)
	if hLen == 0 {
		return 0
	}
	leftMaxes := make([]int, hLen)
	leftMaxes[0] = height[0]
	for i := 1; i < hLen; i++ {
		leftMaxes[i] = max(leftMaxes[i-1], height[i])
	}

	rightMaxes := make([]int, hLen)
	rightMaxes[hLen-1] = height[hLen-1]
	for i := hLen - 2; i >= 0; i-- {
		rightMaxes[i] = max(rightMaxes[i+1], height[i])
	}

	volume := 0
	for i := 1; i < hLen; i++ {
		volume += min(leftMaxes[i], rightMaxes[i]) - height[i]
	}
	return volume
}
