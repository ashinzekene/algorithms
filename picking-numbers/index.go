package algorithms

import (
	"sort"

	. "github.com/ashinzekene/algorithms/utils"
)

func PickingNumbers(a []int) int {
	var maxCount int = 0
	var lowerCount int = 0
	var upperCount int = 0

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

func PickingNumbers1(arr []int) int {
	i, j := 0, 1
	max := 0
	sort.Ints(arr)
	for i < len(arr) {
		for j < len(arr) && Abs(arr[i]-arr[j]) < 2 {
			j++
		}
		max = Max(max, j-i)
		for i < len(arr)-2 && arr[i] == arr[i+1] {
			i++
		}
		i++
		j = i + 1
	}
	return max
}

func Abs(v int) int {
	if v > -1 {
		return v
	}
	return v * -1
}
