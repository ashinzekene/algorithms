package algorithms

import "github.com/ashinzekene/algorithms/utils"

func trap(height []int) int {
	hLen := len(height)
	if hLen == 0 {
		return 0
	}
	leftMaxes := make([]int, hLen)
	leftMaxes[0] = height[0]
	for i := 1; i < hLen; i++ {
		leftMaxes[i] = utils.Max(leftMaxes[i-1], height[i])
	}

	rightMaxes := make([]int, hLen)
	rightMaxes[hLen-1] = height[hLen-1]
	for i := hLen - 2; i >= 0; i-- {
		rightMaxes[i] = utils.Max(rightMaxes[i+1], height[i])
	}

	volume := 0
	for i := 1; i < hLen; i++ {
		volume += utils.Min(leftMaxes[i], rightMaxes[i]) - height[i]
	}
	return volume
}

func trap2(height []int) int {
	length := len(height)
	if length < 3 {
		return 0
	}
	l := 0
	r := length - 1
	l_max := height[0]
	r_max := height[length-1]
	result := 0
	for l < r {
		if height[l] < height[r] {
			if height[l] < l_max {
				result += l_max - height[l]
			} else {
				l_max = height[l]
			}
			l++
		} else {
			if height[r] < r_max {
				result += r_max - height[r]
			} else {
				r_max = height[r]
			}
			r--
		}
	}
	return result
}
