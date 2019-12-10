package algorithms

import "sort"

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		num := nums[i] * -1
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		v1 := i + 1
		v2 := len(nums) - 1
		for v1 < v2 {
			v1Val := nums[v1]
			v2Val := nums[v2]
			if v1Val+v2Val == num {
				result = append(result, []int{nums[i], v1Val, v2Val})
				for v1 < v2 && nums[v1] == nums[v1+1] {
					v1++
				}
				for v1 < v2 && nums[v2] == nums[v2-1] {
					v2--
				}
				v1++
			} else if (v1Val + v2Val) > num {
				v2--
			} else {
				v1++
			}
		}
	}
	return result
}
