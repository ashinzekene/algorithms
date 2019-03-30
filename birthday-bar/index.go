package algorithms

func BirthdayBar(s []int32, d int32, m int32) int32 {
	var result int32 = 0
	var array = make([]int32, 0, 3)
	var sum int32 = 0

	for i := 0; i < len(s); i++ {
		array = []int32{s[i]}
		sum = s[i]
		if sum == d && len(array) == int(m) {
			result++
		}
		for j := i + 1; j < len(s); j++ {
			sum += s[j]
			array = append(array, s[j])
			if sum == d && len(array) == int(m) {
				result++
			}
			if sum > d || len(array) > int(m) {
				// Exit inner loop
			}
		}

	}
	return result
}
