package algorithms

import "math"

func largestRectangle(h []int32) int64 {
	var max int64 = 0
	iStack := make([]int32, 0)
	hStack := make([]int32, 0)
	for i, v := range h {
		stackLen := len(hStack)
		if stackLen == 0 || hStack[stackLen-1] <= v {
			iStack = append(iStack, int32(i))
			hStack = append(hStack, v)
		} else {
			var lastIndex int32
			for len(iStack) > 0 && hStack[len(hStack)-1] > v {
				lastIndex = iStack[len(iStack)-1]
				lastVal := hStack[len(hStack)-1]
				area := lastVal * (int32(i) - lastIndex)
				max = int64(math.Max(float64(max), float64(area)))
				iStack = iStack[:len(iStack)-1]
				hStack = hStack[:len(hStack)-1]
			}
			iStack = append(iStack, lastIndex)
			hStack = append(hStack, v)
		}
	}
	for len(iStack) > 0 {
		lastIndex := iStack[len(iStack)-1]
		lastVal := hStack[len(hStack)-1]
		area := int64(lastVal * (int32(len(h)) - lastIndex))
		max = int64(math.Max(float64(max), float64(area)))
		iStack = iStack[:len(iStack)-1]
		hStack = hStack[:len(hStack)-1]
	}
	return max
}
