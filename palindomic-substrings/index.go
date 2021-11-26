package algorithms

func CountSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s); i++ {
		count += countPalindromes(s, i, i)
		count += countPalindromes(s, i, i+1)
	}
	return count
}

func countPalindromes(s string, l, r int) int {
	count := 0
	for l >= 0 && r < len(s) {
		if s[l] != s[r] {
			return count
		}
		count++
		l--
		r++
	}
	return count
}
