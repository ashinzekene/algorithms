package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_3Sum(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected [][]int
	}{
		{
			[]int{-1, 0, 1, 2, -1, -4},
			[][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			[]int{0, 0, 0},
			[][]int{
				{0, 0, 0},
			},
		},
		{
			[]int{1, -1, -1, 0},
			[][]int{
				{-1, 0, 1},
			},
		},
	}
	for _, tc := range testCases {
		actual := threeSum(tc.Nums)
		assert.Equal(t, tc.Expected, actual)
	}
}
