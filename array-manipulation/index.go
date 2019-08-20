package algorithms

func ArrayManipulation(n int32, queries [][]int32) int64 {
	arr := make([]int64, n)
	var max int64 = 0
	var tsum int64 = 0
	for _, query := range queries {
		arr[query[0]-1] += int64(query[2])
		if query[1] < n {
			arr[query[1]] += -int64(query[2])
		}
	}
	for _, i := range arr {
		tsum += i
		if tsum > max {
			max = tsum
		}
	}
	return max
}
