package algorithms

// https://leetcode.com/problems/number-of-good-ways-to-split-a-string

func NumSplits(s string) int {
	lhMap := make(map[rune]int)
	rhMap := make(map[rune]int)

	for _, c := range s {
		lhMap[c]++
	}

	result := 0
	for _, c := range s {
		if lhMap[c] > 0 {
			lhMap[c]--
		}
		if lhMap[c] == 0 {
			delete(lhMap, c)
		}
		rhMap[c]++
		if len(lhMap) == len(rhMap) {
			result++
		}
	}

	return result
}
