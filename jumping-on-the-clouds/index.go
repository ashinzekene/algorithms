package algorithms

func JumpingOnClouds(c []int32) int32 {
	var counter int32 = 0
	var result int32 = 0
	for counter < int32(len(c)-1) {
		if counter+2 < int32(len(c)) && c[counter+2] == 0 {
			counter++
		}
		counter++
		result++
	}
	return result
}
