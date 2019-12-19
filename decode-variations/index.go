package algorithms

import (
	"strconv"
)

func decodeVariations(S string) int {
	a, b := 0, 0
	if len(S) == 0 {
		return 1
	}
	if S[0] != '0' {
		a = decodeVariations(S[1:])
	}
	if len(S) > 1 {
		val, _ := strconv.Atoi(string(S[:2]))
		if val > 0 && val <= 26 {
			b = decodeVariations(S[2:])
		}
	}
	return a + b
}
