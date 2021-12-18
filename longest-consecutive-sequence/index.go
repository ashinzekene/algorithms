package algorithms

import (
	. "github.com/ashinzekene/algorithms/utils"
)

func LongestConsecutive(nums []int) int {
	numsMap := make(map[int]bool)
	for _, v := range nums {
		numsMap[v] = true
	}

	maxSeqLength := 0
	currentSeqLength := 0
	for num := range numsMap {
		if numsMap[num-1] {
			continue
		}
		currentSeqLength = 1
		currentNum := num
		for numsMap[currentNum+1] {
			currentSeqLength += 1
			currentNum += 1
		}
		maxSeqLength = Max(maxSeqLength, currentSeqLength)
	}
	return maxSeqLength
}
