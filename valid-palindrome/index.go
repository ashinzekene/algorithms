package algorithms

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	l := 0
	r := len(s) - 1
	for l < r {
		lVal := int(s[l])
		rVal := int(s[r])
		if (lVal > 122 || lVal < 48) || (lVal > 57 && lVal < 97) {
			l++
			continue
		}
		if (rVal > 122 || rVal < 48) || (rVal > 57 && rVal < 97) {
			r--
			continue
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
