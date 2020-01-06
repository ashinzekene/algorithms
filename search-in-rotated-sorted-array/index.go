package algorithms

func search(nums []int, target int) int {
	// find pivot
	arrLength := len(nums)
	if arrLength == 0 {
		return -1
	}
	pivot := findPivot(nums, 0, arrLength-1)
	if pivot != -1 && nums[pivot] == target {
		return pivot
	}
	if pivot == -1 {
		pivot = arrLength
	}
	var l, r int
	if target < nums[0] {
		l = pivot + 1
		r = arrLength - 1
	} else {
		l = 0
		r = pivot - 1
	}
	for l <= r {
		m := (l + r) / 2
		if nums[m] == target {
			return m
		} else if target < nums[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return -1
}

// arr[] = {5, 6, 7, 8, 9, 10, 1, 2, 3};
func findPivot(arr []int, l, r int) int {
	if r < l {
		return -1
	}
	if r == l {
		return l
	}
	m := (l + r) / 2
	if l == m && arr[m] > arr[r] {
		return m
	}
	if arr[l] > arr[m] {
		return findPivot(arr, l, m)
	}
	if arr[m] > arr[r] {
		return findPivot(arr, m, r)
	}
	return -1
}
