package algorithms

func searchRange(nums []int, target int) []int {
	i := 0
	j := len(nums) - 1
	result := []int{-1, -1}
	for i <= j {
		if result[0] > -1 && result[1] > -1 {
			return result
		}
		if target == nums[i] {
			result[0] = i
		} else {
			i++
		}
		if target == nums[j] {
			result[1] = j
		} else {
			j--
		}
	}
	return result
}

func searchRange1(nums []int, target int) []int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		m := (j + i) / 2
		val := nums[m]
		if val < target {
			i = m + 1
		} else if val >= target {
			j = m - 1
		}
	}

	if i > len(nums)-1 || nums[i] != target {
		return []int{-1, -1}
	}
	start := i

	i = 0
	j = len(nums) - 1
	for i <= j {
		m := (i + j) / 2
		val := nums[m]
		if val <= target {
			i = m + 1
		} else {
			j = m - 1
		}
	}

	return []int{start, j}
}
