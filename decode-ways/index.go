package algorithms

import "strconv"

func numDecodings(s string) int {
	result := 0
	if s == "" {
		return 1
	}
	if len(s) > 1 {
		num, _ := strconv.Atoi(s[:2])
		if num <= 26 && num > 0 {
			result += numDecodings(s[2:])
		}
	}
	result += numDecodings(s[1:])
	return result
}
