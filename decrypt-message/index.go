package algorithms

func DecryptMessage(word string) string {
	result := ""
	prev := 0
	for i := 0; i < len(word); i++ {
		val := int(word[i])
		if i == 0 {
			result += string(rune(val - 1))
		} else {
			intVal := val - prev
			for intVal < 97 {
				intVal += 26
			}
			result += string(rune(intVal))
		}
		prev = val
	}
	return result
}
