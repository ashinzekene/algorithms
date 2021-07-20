package algorithms

func plusOne(nums []int) []int {
	remainder := 1
	for i := len(nums) - 1; i >= 0; i-- {
		val := nums[i] + remainder
		remainder = val / 10
		nums[i] = val % 10
	}
	if remainder > 0 {
		return append([]int{remainder}, nums...)
	}
	return nums
}
