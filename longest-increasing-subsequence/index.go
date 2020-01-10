package algorithms

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sequenceList := make([]int, 1)
	sequenceList[0] = nums[0]
	for _, v := range nums[1:] {
		isAdded := false
		for i := 0; i < len(sequenceList); i++ {
			if sequenceList[i] >= v {
				sequenceList[i] = v
				isAdded = true
				break
			}
		}
		if !isAdded {
			sequenceList = append(sequenceList, v)
		}
	}
	return len(sequenceList)
}
