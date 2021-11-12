package algorithms

func LengthOfLongestSubstring(s string) int {
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

func Max(a ...int) int {
	max := a[0]
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}
