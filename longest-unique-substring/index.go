package algorithms

import . "github.com/ashinzekene/algorithms/utils"

func lengthOfLongestSubstring(s string) int {
	letterMap := make(map[byte]int)
	pointer := 0
	maxLen := 0
	for i := 0; i < len(s); i++ {
		char := s[i]
		val, present := letterMap[char]
		if present && pointer <= val {
			pointer = letterMap[char] + 1
		}
		letterMap[char] = i
		diff := i - pointer + 1
		maxLen = Max(diff, maxLen)
	}
	return maxLen
}

func lengthOfLongestSubstring1(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	max := 0
	l1 := 0
	l2 := 0
	letterMap := make(map[string]bool)
	for l2 < len(s) {
		if letterMap[string(s[l2])] {
			l1++
			max = Max(max, len(letterMap))
			letterMap = make(map[string]bool)
			letterMap[string(s[l1])] = true
			l2 = l1 + 1
		} else {
			letterMap[string(s[l2])] = true
			l2++
		}
	}
	return Max(max, len(letterMap))
}
