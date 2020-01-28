package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SpiralMatrix(t *testing.T) {
	testCases := []struct {
		Matrix   [][]int
		Expected []int
	}{
		{
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			[][]int{
				[]int{1, 2, 3, 4},
				[]int{5, 6, 7, 8},
				[]int{9, 10, 11, 12},
			},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	for _, tc := range testCases {
		actual := spiralOrder(tc.Matrix)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for %v", tc.Matrix))
	}
}
