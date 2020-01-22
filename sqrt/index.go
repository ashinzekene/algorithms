package algorithms

func mySqrt(x int) int {
	left := 0
	right := x
	for left <= right {
		mid := (left + right) / 2

		if mid*mid > x {
			right = mid - 1
		} else if (mid+1)*(mid+1) <= x {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
