package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MinimumPathSum(t *testing.T) {
	testCases := []struct {
		Grid     [][]int
		Expected int
	}{
		{
			[][]int{
				[]int{0, 0, 0},
				[]int{0, 1, 0},
				[]int{0, 0, 0},
			},
			0,
		},
		{
			[][]int{
				[]int{1, 0},
			},
			1,
		},
		{
			[][]int{
				[]int{0, 0, 0, 0},
				[]int{0, 1, 2, 3},
				[]int{0, 0, 5, 0},
			},
			3,
		},
		{
			[][]int{
				[]int{1, 1, 3},
				[]int{3, 2, 1},
				[]int{1, 2, 1},
			},
			6,
		},
		{
			[][]int{
				[]int{1, 2, 5},
				[]int{3, 2, 1},
			},
			6,
		},
	}
	for _, tc := range testCases {
		actual := minPathSum(tc.Grid)
		assert.Equal(t, tc.Expected, actual)
	}
}
