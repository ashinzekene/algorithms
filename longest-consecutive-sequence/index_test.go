package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LongestConsecutive(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			5,
		},
		{
			[]int{2, 3, 4},
			3,
		},
		{
			[]int{2},
			1,
		},
		{
			[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			9,
		},
		{
			[]int{5, 1, 2, 3, 4},
			5,
		},
		{
			[]int{100, 4, 200, 1, 3, 2},
			4,
		},
	}
	for _, tc := range testCases {
		actual := LongestConsecutive(tc.Nums)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for Nums: %v", tc.Nums))
	}
}
