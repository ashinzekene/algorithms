package algorithms

func singleNumber(nums []int) int {
	numsMap := make(map[int]int)

	for _, v := range nums {
		numsMap[v]++
	}

	for integer, count := range numsMap {
		if count == 1 {
			return integer
		}
	}
	return -1
}
