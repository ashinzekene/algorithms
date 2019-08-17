package algorithms

func CountingSort(arr []int32) []int32 {
	result := make([]int32, len(arr))
	var cummualative int32 = 0
	var max int32 = 0
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	sorterArr := make([]int32, int(max))
	for _, val := range arr {
		sorterArr[val]++
	}
	for key, val := range sorterArr {
		sorterArr[key] += cummualative
		cummualative += val
	}
	// for _, val := range arr {

	// }
	return result
}

// [4,2,3,4,5,1,2,8,5]
// [1 => ,2 =>,3 => ,4 =>,5 =>,8=> ]
