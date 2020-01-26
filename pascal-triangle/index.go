package algorithms

func generate(numRows int) [][]int {
	result := make([][]int, 0)
	for i := 0; i < numRows; i++ {
		if i == 0 {
			result = append(result, []int{1})
		} else if i == 1 {
			result = append(result, []int{1, 1})
		} else {
			lastSlice := result[len(result)-1]
			newSlice := make([]int, len(lastSlice)+1)
			newSlice[0] = 1
			for i := 1; i < len(lastSlice); i++ {
				newSlice[i] = lastSlice[i-1] + lastSlice[i]
			}
			newSlice[len(newSlice)-1] = 1
			result = append(result, newSlice)
		}
	}
	return result
}
