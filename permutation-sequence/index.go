package algorithms

import "strconv"

func getPermutation(n int, k int) string {
	arr := make([]bool, n+1)
	result := ""
	k--
	for ; n > 1; n-- {
		t := permutate(n - 1)
		quo := k / t
		k = k % t
		val := getVal(arr, quo)
		result += strconv.Itoa(val)
	}
	return result + strconv.Itoa(getVal(arr, 0))
}

func getVal(arr []bool, remainder int) int {
	k := -1
	for i := 1; i < len(arr); i++ {
		if !arr[i] {
			k++
		}
		if k == remainder {
			arr[i] = true
			return i
		}
	}
	return 0
}

func permutate(n int) int {
	result := 1
	for ; n > 1; n-- {
		result *= n
	}
	return result
}
