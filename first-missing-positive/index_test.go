package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_firstMissingPositive(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected int
	}{
		{
			[]int{2},
			1,
		},
		{
			[]int{1, 2, 0},
			3,
		},
		{
			[]int{3, 4, -1, 1},
			2,
		},
		{
			[]int{7, 8, 9, 11, 12},
			1,
		},
		{
			[]int{2, 5, 10, 9, 1, 3, 4},
			6,
		},
		{
			[]int{-4},
			1,
		},
	}
	for _, tc := range testCases {
		actual := firstMissingPositive1(tc.Nums)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for Digit %d", tc.Nums))
	}
}
