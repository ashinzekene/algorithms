package algorithms

func permuteUnique(nums []int) [][]int {
	result := [][]int{}
	usedMap := make([]bool, len(nums))
	perm(&result, usedMap, nums, []int{})
	return result
}

func perm(result *[][]int, usedMap []bool, nums, current []int) {
	if len(current) == len(nums) {
		newList := make([]int, len(current))
		copy(newList, current)
		*result = append(*result, newList)
		return
	}
	loopUsed := make(map[int]bool)
	for i, v := range nums {
		if usedMap[i] || loopUsed[v] {
			continue
		}
		usedMap[i] = true
		loopUsed[v] = true
		perm(result, usedMap, nums, append(current, v))
		usedMap[i] = false
	}
}
