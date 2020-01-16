package algorithms

func firstMissingPositive(nums []int) int {
	lenNums := len(nums)
	result := 1
	if lenNums == 0 {
		return result
	}
	length := lenNums + 1
	maxVal := nums[0]
	for _, val := range nums {
		if val > maxVal {
			maxVal = val
		}
	}
	if maxVal < 1 {
		return result
	}
	if maxVal < length {
		length = maxVal + 1
	}
	numsMap := make([]int, length)

	for _, val := range nums {
		if val > 0 && val < length {
			numsMap[val]++
		}
	}

	for i := 1; i < length; i++ {
		if numsMap[i] == 0 {
			return i
		}
	}
	return maxVal + 1
}
