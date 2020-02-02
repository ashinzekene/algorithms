package algorithms

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	lastGood := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i]+i >= lastGood {
			lastGood = i
		}
	}
	return lastGood == 0
}
