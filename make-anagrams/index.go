package algorithms

import "strings"

func MakeAnagram(a string, b string) int32 {
	var count int32 = 0
	c := b
	for _, r := range a {
		if strings.Contains(c, string(r)) {
			c = strings.Replace(c, string(r), "", 1)
		} else {
			count++
		}
	}
	return count + int32(len(c))
}
