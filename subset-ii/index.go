package algorithms

import "sort"

func subsetsWithDup(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums)
	for length := 0; length <= len(nums); length++ {
		addSubSet(&result, []int{}, nums, length, 0)
	}
	return result
}

func addSubSet(result *[][]int, currentSubSet []int, numsLeft []int, length, start int) {
	if len(currentSubSet) == length {
		subset := make([]int, len(currentSubSet))
		copy(subset, currentSubSet)
		*result = append(*result, subset)
		return
	}
	for i := start; i < len(numsLeft); i++ {
		if i > start && numsLeft[i] == numsLeft[i-1] {
			continue
		}
		addSubSet(result, append(currentSubSet, numsLeft[i]), numsLeft, length, i+1)
	}
}
