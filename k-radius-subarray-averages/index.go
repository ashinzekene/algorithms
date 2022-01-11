package algorithms

func getAverages(nums []int, k int) []int {
	l := len(nums)
	cummulativeNums := make([]int, l+1)
	result := make([]int, l)
	for i, v := range nums {
		cummulativeNums[i+1] = cummulativeNums[i] + v
	}
	for i, _ := range nums {
		if i-k < 0 || i+k >= l {
			result[i] = -1
			continue
		}
		sum := cummulativeNums[i+k+1] - cummulativeNums[i-k]
		// fmt.Println(i-k, i+k, sum)
		result[i] = sum / (k*2 + 1)
	}
	return result
}
