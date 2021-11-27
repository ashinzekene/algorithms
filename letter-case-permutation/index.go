package algorithms

import "strings"

func LetterCasePermutation(s string) []string {
	result := make([]string, 0)
	getCasePermutation(&result, s, "", 0)
	return result
}

func getCasePermutation(result *[]string, s, currentS string, i int) {
	if i == len(s) {
		*result = append(*result, currentS)
		return
	}
	current := s[i : i+1]
	lowerCase := strings.ToLower(current)
	upperCase := strings.ToUpper(current)
	if lowerCase != upperCase {
		getCasePermutation(result, s, currentS+upperCase, i+1)
		getCasePermutation(result, s, currentS+lowerCase, i+1)
	} else {
		getCasePermutation(result, s, currentS+current, i+1)
	}
}
