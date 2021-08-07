package algorithms

import (
	"sort"

	. "github.com/ashinzekene/algorithms/utils"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; {
		v := nums[i]
		j := i + 1
		k := len(nums) - 1
		for j < k {
			val1 := nums[j]
			val2 := nums[k]
			sum := v + val1 + val2
			diff := Abs(sum - target)
			prevDiff := Abs(result - target)
			if diff < prevDiff {
				result = sum
			}
			if result == target {
				return result
			} else if sum < target {
				j++
			} else {
				k--
			}
		}
		i++
	}
	return result
}
