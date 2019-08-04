package algorithms

func EqualizeArray(arr []int32) int32 {
	l := len(arr)
	var max int32
	strMap := make(map[int32]int32)
	for _, v := range arr {
		strMap[v]++
	}
	for _, v := range strMap {
		if v > max {
			max = v
		}
	}
	return int32(l) - max
}
