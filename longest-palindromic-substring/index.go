package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils"
)

func LongestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start := 0
	length := 0
	for i := 0; i < len(s); i++ {
		l1 := getlengthPalindromeFrom(s, i, i)
		l2 := getlengthPalindromeFrom(s, i, i+1)
		l := Max(l1, l2)
		if l > length {
			start = i - (l-1)/2
			length = l
		}
	}
	return s[start : start+length]
}

func getlengthPalindromeFrom(s string, l, r int) int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return r - l - 1
}
