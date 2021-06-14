package algorithms

import "math"

func firstMissingPositive(nums []int) int {
	max := 0
	for _, v := range nums {
		if v > max {
			max = v
		}
	}

	arrMap := make([]bool, max+1)
	arrMap[0] = true

	for _, v := range nums {
		if v > 0 {
			arrMap[v] = true
		}
	}

	for i := range arrMap {
		if !arrMap[i] {
			return i
		}
	}

	return max + 1
}

func firstMissingPositive1(nums []int) int {
	l := len(nums)
	for i := range nums {
		if nums[i] <= 0 || nums[i] > l {
			nums[i] = len(nums)
		}
	}

	for _, num := range nums {
		num = int(math.Abs(float64(num)))
		if num == len(nums) {
			continue
		}
		numIndex := num - 1
		if nums[numIndex] < 0 {
			continue
		}
		nums[numIndex] *= -1
	}

	for i, num := range nums {
		if num >= 0 {
			return i + 1
		}
	}

	return len(nums)
}
