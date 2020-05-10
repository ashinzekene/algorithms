package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UniquPathsII(t *testing.T) {
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
			2,
		},
		{
			[][]int{
				[]int{1, 0},
			},
			0,
		},
		{
			[][]int{
				[]int{0, 0, 0, 0},
				[]int{0, 1, 0, 0},
				[]int{0, 0, 0, 0},
			},
			4,
		},
		{
			[][]int{
				[]int{0, 0, 0, 0},
				[]int{0, 1, 0, 0},
				[]int{0, 0, 0, 1},
				[]int{0, 0, 0, 0},
			},
			4,
		},
	}
	for _, tc := range testCases {
		actual := uniquePathsWithObstacles(tc.Grid)
		assert.Equal(t, tc.Expected, actual)
	}
}
