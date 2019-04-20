package algorithms

func MinimumBribes(arr []int32) interface{} {
	bribes := 0
	for i, val := range arr {
		if int(val) > i+3 {
			return "Too chaotic"
		}
		no := noOfGreater(val, arr[i:])
		if no == -1 {
			return "Too chaotic"
		} else {
			bribes += no
		}
	}
	return bribes
}

func noOfGreater(n int32, arr []int32) int {
	count := 0
	for _, x := range arr {
		if n > x {
			count++
			if count > 2 {
				return -1
			}
		}
	}
	return count
}
