package algorithms

func FindDigits(n int32) int32 {
	var divisor int32 = 10
	var count int32 = 0
	x := n
	for x > 0 {
		digit := x % divisor
		if digit != 0 && n%digit == 0 {
			count++
		}
		x /= 10
	}
	return count
}
