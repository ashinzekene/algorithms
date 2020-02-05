package algorithms

import "strconv"

func numDecodings(s string) int {
	result := make([]int, len(s)+1)
	result[0] = 1
	for i := 1; i <= len(s); i++ {
		val, _ := strconv.Atoi(string(s[i-1]))
		if val > 0 {
			result[i] += result[i-1]
		}
		if i == 1 {
			continue
		}
		prev, _ := strconv.Atoi(string(s[i-2]))
		if prev == 1 || (prev == 2 && val <= 6) {
			result[i] += result[i-2]
		}
	}
	return result[len(result)-1]
}
