package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findPeakGrid(t *testing.T) {
	testCases := []struct {
		Matrix   [][]int
		Expected []int
	}{
		{
			[][]int{
				{1, 4},
				{3, 2},
			},
			[]int{0, 1},
		},
		{
			[][]int{
				{10, 20, 15},
				{21, 30, 14},
				{7, 16, 32},
			},
			[]int{1, 1},
		},
		{
			[][]int{
				{7, 2, 3, 1, 2},
				{6, 5, 4, 2, 1},
			},
			[]int{0, 0},
		},
	}
	for _, tc := range testCases {
		actual := findPeakGrid(tc.Matrix)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for %v", tc.Matrix))
	}
}
