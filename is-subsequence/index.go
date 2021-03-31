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
