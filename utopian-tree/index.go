package algorithms

func UtopianTree(n int32) int32 {
	var height int32 = 1
	for x := 1; int32(x) <= n; x++ {
		if x%2 == 1 {
			height *= 2
		} else {
			height++
		}
	}
	return height
}
