package algorithms

func AlternatingCharacters(s string) int32 {
	var lastType string
	var count int32 = 0
	for i, v := range s {
		if i == 0 {
			lastType = string(v)
		} else {
			if string(v) == lastType {
				count++
			} else {
				lastType = string(v)
			}
		}
	}
	return count
}
