package algorithms

import "fmt"

func countAndSay(n int) string {
	str := "1"
	for i := 1; i < n; i++ {
		newStr := ""
		start := 0
		end := 0
		for end <= len(str) {
			if end == len(str) || str[start] != str[end] {
				newStr += fmt.Sprintf("%d%s", end-start, string(str[start]))
				start = end
			}
			end++
		}
		str = newStr
	}
	return str
}
