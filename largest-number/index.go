package algorithms

import (
	"sort"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/largest-number/

func largestNumber(nums []int) string {
	newNums := make([]string, len(nums))
	for i, v := range nums {
		newNums[i] = strconv.Itoa(v)
	}
	sort.Slice(newNums, func(i, j int) bool {
		str1 := newNums[i]
		str2 := newNums[j]
		return str1+str2 > str2+str1
	})

	num := strings.Join(newNums, "")
	for len(num) > 1 && string(num[0]) == "0" {
		num = num[1:]
	}
	return num
}
