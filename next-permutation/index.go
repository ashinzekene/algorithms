package algorithms

func nextPermutation(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	avail := make(map[int]int)
	l := 0
	avail[nums[len(nums)-1]] = len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i+1] > nums[i] {
			val := nums[i] + 1
			for avail[val] <= 0 {
				val++
			}
			l = i + 1
			temp := nums[i]
			nums[i] = val
			nums[avail[val]] = temp
			break
		} else if avail[nums[i]] == 0 {
			avail[nums[i]] = i
		}
	}

	r := len(nums) - 1
	for l < r {
		temp := nums[l]
		nums[l] = nums[r]
		nums[r] = temp
		l++
		r--
	}
	return nums
}
