package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MinimumPathSum(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Target   int
		Expected int
	}{
		{
			[]int{2, 3, 1, 2, 4, 3},
			7,
			2,
		},
		{
			[]int{1, 4, 4},
			4,
			1,
		},
		{
			[]int{1, 1, 1, 1, 1, 1, 1, 1},
			11,
			0,
		},
		{
			[]int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12},
			213,
			8,
		},
	}
	for _, tc := range testCases {
		actual := minSubArrayLen(tc.Target, tc.Nums)
		assert.Equal(t, tc.Expected, actual)
	}
}
