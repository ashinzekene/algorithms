package algorithms

func convert(s string, numRows int) string {
	rowSlice := make([]string, numRows)

	if numRows == 1 {
		return s
	}

	j := 0
	addJ := false
	for _, char := range s {
		rowSlice[j] += string(char)
		if j == numRows-1 || j == 0 {
			addJ = !addJ
		}
		if addJ {
			j++
		} else {
			j--
		}
	}
	result := ""
	for _, row := range rowSlice {
		result += row
	}

	return result
}
