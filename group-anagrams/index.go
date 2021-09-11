package algorithms

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	stringMapIndex := make(map[string]int)
	for _, s := range strs {
		sortedStr := sortIt(s)
		if index, ok := stringMapIndex[sortedStr]; ok {
			result[index] = append(result[index], s)
		} else {
			stringMapIndex[sortedStr] = len(result)
			result = append(result, []string{s})
		}
	}

	return result
}

func sortIt(s string) string {
	strRune := []rune(s)
	sort.Slice(strRune, func(i, j int) bool {
		return strRune[i] < strRune[j]
	})
	return string(strRune)
}

func groupAnagrams1(strs []string) [][]string {
	result := make([][]string, 0)
	alphas := make(map[[26]byte]int)

	for _, str := range strs {
		alpha := getAlpha(str)
		if index, exists := alphas[alpha]; exists {
			result[index] = append(result[index], str)
		} else {
			alphas[alpha] = len(result)
			result = append(result, []string{str})
		}
	}
	return result
}

func getAlpha(s string) [26]byte {
	alpha := [26]byte{}
	for _, char := range s {
		alpha[char-97]++
	}
	return alpha
}
