package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LeftRotation(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected int
	}{
		{
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			6,
		},
		{
			[]int{4, -1, 2, 1},
			6,
		},
		{
			[]int{3, 5, -1, 4, 2, -2, -3, 3},
			13,
		},
	}
	for _, tc := range testCases {
		actual := maxSubArray(tc.Nums)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Nums))
	}
}
