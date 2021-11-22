package algorithms

func Partition(s string) [][]string {
	palindromeSubstrings := make(map[int]map[int]bool)
	result := make([][]string, 0)
	if s == "" {
		return result
	}
	dfs(palindromeSubstrings, 0, s, []string{}, &result)
	return result
}

func dfs(dp map[int]map[int]bool, start int, s string, currentPartition []string, result *[][]string) {
	if start >= len(s) {
		p := make([]string, len(currentPartition))
		copy(p, currentPartition)
		*result = append(*result, p)
		return
	}
	for end := start; end < len(s); end++ {
		if s[end] == s[start] && (dp[start+1][end-1] || (end-start <= 2)) {
			if dp[start] == nil {
				dp[start] = make(map[int]bool)
			}
			dp[start][end] = true
			dfs(dp, end+1, s, append(currentPartition, s[start:end+1]), result)
		}
	}
}

func Partition1(s string) [][]string {
	palindromeSubstrings := make(map[int]map[int]bool)
	result := make([][]string, 0)
	if s == "" {
		return result
	}
	for i := 0; i < len(s); i++ {
		getlengthPalindromeFrom(palindromeSubstrings, s, i, i)
		getlengthPalindromeFrom(palindromeSubstrings, s, i, i+1)
	}
	getPartitions(&result, palindromeSubstrings, 0, s, []string{})
	return result
}

func getPartitions(result *[][]string, palindromeSubstrings map[int]map[int]bool, start int, s string, currentPartition []string) {
	if start >= len(s) {
		p := make([]string, len(currentPartition))
		copy(p, currentPartition)
		*result = append(*result, p)
		return
	}
	for i := 0; i < len(s); i++ {
		end := i + start
		if v := palindromeSubstrings[start]; v != nil && v[end] {
			getPartitions(result, palindromeSubstrings, end+1, s, append(currentPartition, s[start:end+1]))
		}
	}
}

func getlengthPalindromeFrom(palindromeSubstrings map[int]map[int]bool, s string, l, r int) int {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		if palindromeSubstrings[l] == nil {
			palindromeSubstrings[l] = make(map[int]bool)
		}
		palindromeSubstrings[l][r] = true
		l--
		r++
	}
	return r - l - 1
}
