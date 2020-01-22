package algorithms

func plusOne(digits []int) []int {
	digits_length := len(digits)
	if digits_length == 0 {
		return digits
	}
	i := digits_length - 1
	for i > 0 && digits[i] == 9 {
		digits[i] = 0
		i--
	}
	if i == 0 && digits[i] == 9 {
		digits[0] = 0
		digits = append([]int{1}, digits...)
	} else if i >= 0 {
		digits[i]++
	}
	return digits
}
