package algorithms

func GetMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	var max int32 = -1
	for _, x := range keyboards {
		for _, y := range drives {
			sum := x + y
			if sum > int32(max) && sum <= b {
				max = sum
			}
		}
	}
	return max
}
