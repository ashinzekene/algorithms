package algorithms

import "sort"

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; {
		v := nums[i]
		j := i + 1
		k := len(nums) - 1
		for j < k {
			val1 := nums[j]
			val2 := nums[k]
			sum := v + val1 + val2
			if sum == 0 {
				result = append(result, []int{v, val1, val2})
				j++
				for j < len(nums) && nums[j] == nums[j-1] {
					j++
				}
				k--
				for k > 0 && nums[k] == nums[k+1] {
					k--
				}
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
		i++
		for i < len(nums) && nums[i] == nums[i-1] {
			i++
		}
	}
	return result
}
