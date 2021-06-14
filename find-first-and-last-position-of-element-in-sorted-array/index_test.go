package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchRange(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Target   int
		Expected []int
	}{
		{
			[]int{5, 7, 7, 8, 8, 10},
			8,
			[]int{3, 4},
		},
		{
			[]int{1, 2, 3, 5, 6, 7, 8, 8, 9, 10},
			10,
			[]int{9, 9},
		},
		{
			[]int{1, 2, 3, 5, 6, 7, 8, 8, 9, 10},
			8,
			[]int{6, 7},
		},
		{
			[]int{1, 2, 3, 5, 6, 7, 8, 8, 9, 10},
			2,
			[]int{1, 1},
		},
		{
			[]int{1, 2, 3, 5, 6, 7, 8, 8, 9, 10},
			6,
			[]int{4, 4},
		},
		{
			[]int{1, 2, 3, 5, 6, 7, 8, 8, 9, 10},
			4,
			[]int{-1, -1},
		},
		{
			[]int{1, 1},
			4,
			[]int{-1, -1},
		},
		{
			[]int{1, 1},
			1,
			[]int{0, 1},
		},
	}
	for _, tc := range testCases {
		actual := searchRange1(tc.Nums, tc.Target)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for Digit %d", tc.Nums))
	}
}
