package algorithms

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}
	left := getExtreme(nums, target, true)
	if left == len(nums) || nums[left] != target {
		return result
	}
	result[0] = left
	result[1] = getExtreme(nums, target, false) - 1
	return result
}

func getExtreme(nums []int, target int, left bool) int {
	l := 0
	r := len(nums)
	for l < r {
		m := (l + r) / 2
		if nums[m] > target || (nums[m] == target && left) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
