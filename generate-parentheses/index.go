package algorithms

func generateParenthesis(n int) []string {
	result := []string{}
	if n == 0 {
		return result
	}
	addParenthesis("", &result, n-1, 0, 0)
	return result
}

func addParenthesis(str string, result *[]string, n, l, r int) {
	if l == n && r == n {
		*result = append(*result, "("+str+")")
		return
	}
	if l <= r {
		addParenthesis(str+")", result, n, l+1, r)
	}
	if r < n {
		addParenthesis(str+"(", result, n, l, r+1)
	}
}
