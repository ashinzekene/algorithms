package algorithms

func checkMagazine(magazine []string, note []string) string {
	words := make(map[string]int32)

	for _, w := range magazine {
		if _, ok := words[w]; ok {
			words[w]++
		} else {
			words[w] = 1
		}
	}

	for _, w := range note {
		if val, ok := words[w]; !ok || val == 0 {
			return "No"
		} else {
			words[w]--
		}
	}
	return "Yes"
}
