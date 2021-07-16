package algorithms

func wordBreak(s string, wordDict []string) bool {
	sLength := len(s)
	dp := make([]bool, sLength+1)

	dp[0] = true
	for i := range dp {
		for _, word := range wordDict {
			wordLength := len(word)
			start := i - wordLength
			if start >= 0 && dp[start] && word == s[start:i] {
				dp[i] = true
			}
		}
	}
	return dp[sLength]
}
