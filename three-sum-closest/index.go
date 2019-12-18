package algorithms

import (
	"math"
	"sort"
)

func resolveSum(sum int, closestSum int, target int) int {
	diff := math.Abs(float64(sum - target))
	if math.Abs(float64(closestSum-target)) > diff {
		return sum
	}
	return closestSum
}

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	closestSum := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		v1 := nums[i]
		start := i + 1
		end := len(nums) - 1
		for start < end {
			v2 := nums[start]
			v3 := nums[end]
			closestSum = resolveSum(v1+v2+v3, closestSum, target)
			if v1+v2+v3 > target {
				end--
			} else {
				start++
			}
		}
	}
	return closestSum
}
