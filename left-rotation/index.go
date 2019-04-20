package algorithms

func LeftRotation(arr []int32, n int32) []int32 {
	length := len(arr)
	rotations := int(n) % length
	result := make([]int32, length)
	for i := 0; i < length; i++ {
		result[(length-rotations+i)%length] = arr[i]
	}
	return result
}
