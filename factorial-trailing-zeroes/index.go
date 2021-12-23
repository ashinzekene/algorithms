package algorithms

// no of trailing zeros is no of ocurrences of products that result in trailing zeros
// Eg 5! has 1 trailing zero because 2*5 has a trailing zero
// Eg 11! has 2 trailing zero because 2*5 has a trailing zero and 1 *10 has another trailing zero
// Eg 26! has 6 trailing zero because 2*5, 10, 12*15, 20, 12*25, 4*25

func trailingZeroes(n int) int {
	result := 0
	for i := 5; i < n; i *= 5 {
		result += n / i
	}
	return result
}
