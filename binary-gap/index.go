package algorithms

import (
	"strconv"

	. "github.com/ashinzekene/algorithms/utils"
)

func binaryGap(N int) int {
	base10 := strconv.FormatInt(int64(N), 2)
	start := -1
	for i, v := range base10 {
		if v == '1' {
			start = i
			break
		}
	}
	end := start + 1
	max := 0
	current := 0
	for ; end < len(base10); end++ {
		v := base10[end]
		if v == '0' {
			current++
		} else {
			max = Max(current, max)
			current = 0
			start = end //nolint
		}
	}
	return max
}
