package algorithms

func isPalindrome(x int) bool {
	original := x
	reversed := 0
	if original < 0 {
		return false
	}
	for x > 0 {
		reversed = (reversed * 10) + (x % 10)
		x /= 10
	}
	return reversed == original
}
