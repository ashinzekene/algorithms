package algorithms

func isSubsequence(s string, t string) bool {
	i := 0
	j := 0
	for i < len(s) && j < len(t) {
		iVal := s[i]
		jVal := t[j]
		if iVal == jVal {
			i++
		}
		j++
	}
	return i == len(s)
}

func isSubsequence1(s string, t string) bool {
	if s == "" {
		return true
	}
	for i, v := range t {
		if string(v) == string(s[0]) {
			return isSubsequence(s[1:], t[i+1:])
		}
	}
	return false
}
