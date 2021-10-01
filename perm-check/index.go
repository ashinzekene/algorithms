package algorithms

func PermCheck(A []int) int {
	count := 0
	presentMap := make([]bool, len(A)+1)
	for _, val := range A {
		if val > len(A) {
			return 0
		}
		if !presentMap[val] {
			presentMap[val] = true
			count++
		}
	}
	if count == len(A) {
		return 1
	}
	return 0
}
