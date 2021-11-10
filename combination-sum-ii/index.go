package algorithms

import "sort"

func CombinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(candidates)
	checkSum(candidates, []int{}, 0, 0, target, &result)
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
	for j := i; j < len(candidates) && (sum+candidates[j]) <= target; j++ {
		if j > i && candidates[j] == candidates[j-1] {
			continue
		}
		checkSum(candidates, append(slice, candidates[j]), j+1, sum+candidates[j], target, result)
	}
}
