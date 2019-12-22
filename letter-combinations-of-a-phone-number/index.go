package algorithms

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	mapping := map[byte]string{
		'0': "",
		'1': "",
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	result := []string{}
	combinations(digits, "", 0, &result, mapping)
	return result
}

func combinations(
	digits, str string, digitIndex int, result *[]string, mapping map[byte]string,
) {
	if digitIndex == len(digits) {
		*result = append(*result, str)
		return
	} else {
		key := digits[digitIndex]
		for _, i := range mapping[key] {
			combinations(digits, str+string(i), digitIndex+1, result, mapping)
		}
	}
}
