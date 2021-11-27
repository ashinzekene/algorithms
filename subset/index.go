package algorithms

func Subsets(nums []int) [][]int {
	result := [][]int{}
	used := make(map[int]bool)
	for length := 0; length <= len(nums); length++ {
		addSubSet(&result, []int{}, used, nums, length, 0)
	}
	return result
}

func addSubSet(result *[][]int, currentSubSet []int, used map[int]bool, numsLeft []int, length, start int) {
	if len(currentSubSet) == length {
		subset := make([]int, len(currentSubSet))
		copy(subset, currentSubSet)
		*result = append(*result, subset)
		return
	}
	for i := start; i < len(numsLeft); i++ {
		if !used[i] {
			used[i] = true
			addSubSet(result, append(currentSubSet, numsLeft[i]), used, numsLeft, length, i+1)
			used[i] = false
		}
	}
}
