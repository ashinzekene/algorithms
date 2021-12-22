package algorithms

func wiggleMaxLength(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	length := 1
	lastDiff := 0
	lastNum := nums[0]
	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if isSameSigns(lastDiff, diff) || diff == 0 {
			lastDiff = nums[i] - lastNum
		} else {
			length++
			lastDiff = diff
			lastNum = nums[i-1]
		}
	}
	return length
}

func isSameSigns(lastDiff, diff int) bool {
	return lastDiff > 0 && diff > 0 || lastDiff < 0 && diff < 0
}
