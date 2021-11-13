package algorithms

func DailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := make([][2]int, 0)
	for i, temp := range temperatures {
		if len(stack) == 0 {
			stack = append(stack, [2]int{i, temp})
		} else {
			lastItem := stack[len(stack)-1]
			for len(stack) > 0 && lastItem[1] < temp {
				result[lastItem[0]] = i - lastItem[0]
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					lastItem = stack[len(stack)-1]
				}
			}
			stack = append(stack, [2]int{i, temp})
		}
	}
	return result
}
