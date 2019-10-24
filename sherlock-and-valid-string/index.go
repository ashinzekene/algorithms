package algorithms

func isValid(s string) string {
	var val int32 = 0
	clean := true
	strMap := make(map[rune]int32)
	for _, char := range s {
		strMap[char]++
	}
	for _, count := range strMap {
		if val == 0 {
			val = count
		}
		if count != val {
			if !clean {
				return "NO"
			}
			clean = false
		}
	}
	return "YES"
}
