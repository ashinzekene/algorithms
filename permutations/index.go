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

func permute1(arr []int) [][]int {
	result := make([][]int, 0)
	addPermutations(make([]int, 0), arr, &result)
	return result
}

func addPermutations(added, available []int, result *[][]int) {
	if len(available) == 0 {
		*result = append(*result, added)
		return
	}
	for i := range available {
		current := available[i]
		newAvailable := make([]int, len(available))
		copy(newAvailable, available)
		newAvailable = append(newAvailable[:i], newAvailable[i+1:]...)
		addPermutations(append(added, current), newAvailable, result)
	}
}
