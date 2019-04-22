package algorithms

func PickingNumbers(a []int32) int32 {
	var maxCount int32 = 0
	var lowerCount int32 = 0
	var upperCount int32 = 0

	for x := 0; x < len(a); x++ {
		for i := 0; i < len(a); i++ {
			if a[x]-a[i] == 1 {
				lowerCount++
			} else if a[i]-a[x] == 1 {
				upperCount++
			} else if a[x] == a[i] {
				upperCount++
				lowerCount++
			}
		}
		if upperCount > maxCount {
			maxCount = upperCount
		}
		if lowerCount > maxCount {
			maxCount = lowerCount
		}
		upperCount = 0
		lowerCount = 0
	}
	return maxCount
}
