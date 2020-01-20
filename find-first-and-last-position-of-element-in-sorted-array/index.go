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
