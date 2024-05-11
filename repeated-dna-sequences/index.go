package algorithms

func findRepeatedDnaSequences(s string) []string {
	patternCount := make(map[string]int)
	result := make([]string, 0)
	if len(s) < 10 {
		return result
	}
	for i := range s[:len(s)-9] {
		pattern := s[i : i+10]
		patternCount[pattern]++
	}
	for pattern, count := range patternCount {
		if count > 1 {
			result = append(result, pattern)
		}
	}
	return result
}
