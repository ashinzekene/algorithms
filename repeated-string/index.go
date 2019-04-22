package algorithms

import "strings"

func RepeatedString(s string, n int64) int64 {
	parts := n / int64(len(s))
	countPerPart := int64(strings.Count(s, "a"))
	restLength := n % int64(len(s))
	remainder := int64(strings.Count(s[:restLength], "a"))
	return parts*countPerPart + remainder
}
