package algorithms

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) == 0 {
		return result
	}
	for i, val := range nums {
		present := make(map[int]bool)
		present[i] = true
		addToResult(nums, &result, []int{val}, present, i)
	}
	return result
}

func addToResult(nums []int, result *[][]int, current []int, present map[int]bool, i int) {
	if len(current) == len(nums) {
		*result = append(*result, current)
		return
	}
	for index, val := range nums {
		if present[index] {
			continue
		}
		present[index] = true
		addToResult(nums, result, append(current, val), present, index)
		present[index] = false
	}
}