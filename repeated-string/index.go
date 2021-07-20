package algorithms

import "strings"

func RepeatedString(s string, n int64) int64 {
	parts := n / int64(len(s))
	countPerPart := int64(strings.Count(s, "a"))
	restLength := n % int64(len(s))
	remainder := int64(strings.Count(s[:restLength], "a"))
	return parts*countPerPart + remainder
}

func RepeatedString1(s string, n int64) int64 {
	l := len(s)
	repeats := n / int64(l)
	var totalAOcurrencePerString int64 = 0
	for _, v := range s {
		if v == 'a' {
			totalAOcurrencePerString++
		}
	}
	totalAOcurrencePerString *= repeats
	for i := 0; int64(i) < (n % int64(l)); i++ {
		if s[i] == 'a' {
			totalAOcurrencePerString++
		}
	}
	return totalAOcurrencePerString
}
