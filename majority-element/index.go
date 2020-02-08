package algorithms

func majorityElement(nums []int) int {
	numMap := make(map[int]int)
	for _, v := range nums {
		numMap[v]++
	}
	maxCount := 0
	maxInt := nums[0]
	for integer, count := range numMap {
		if count > maxCount {
			maxCount = count
			maxInt = integer
		}
	}
	return maxInt
}
