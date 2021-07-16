package algorithms

func wordBreak(s string, wordDict []string) []string {
	results := make([]string, 0)
	findWords(s, "", wordDict, &results)
	return results
}

func findWords(s, currentS string, wordDict []string, results *[]string) {
	if s == "" {
		*results = append(*results, currentS[1:])
	}
	for _, word := range wordDict {
		l := len(word)
		if l <= len(s) && word == s[:l] {
			findWords(s[l:], currentS+" "+word, wordDict, results)
		}
	}
}
