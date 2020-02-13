package algorithms

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(candidates)
	for i, v := range candidates {
		checkSum(candidates, []int{v}, i, v, target, &result)
	}
	return result
}

func checkSum(candidates, slice []int, i, sum, target int, result *[][]int) {
	if sum > target {
		return
	}
	if sum == target {
		*result = append(*result, slice)
		return
	}
	if i == len(candidates) {
		return
	}
	checkSum(candidates, append(slice, candidates[i]), i, sum+candidates[i], target, result)
	for j := i + 1; j < len(candidates) && (sum+candidates[j]) <= target; j++ {
		checkSum(candidates, append(append([]int{}, slice...), candidates[j]), j, sum+candidates[j], target, result)
	}
}
