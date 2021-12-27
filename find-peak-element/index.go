package algorithms

func findPeakElement(nums []int) int {
	r := len(nums) - 1
	l := 0
	for r > l {
		m := l + (r-l)/2
		if nums[m+1] > nums[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
