package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RotateImage(t *testing.T) {
	testCases := []struct {
		Matrix   [][]int
		Expected [][]int
	}{
		{
			[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			[][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			[][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			[][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
		{
			[][]int{
				{1, 2},
				{3, 4},
			},
			[][]int{
				{3, 1},
				{4, 2},
			},
		},
	}
	for _, tc := range testCases {
		rotate(tc.Matrix)
		assert.Equal(t, tc.Expected, tc.Matrix, fmt.Sprintf("Did not work for %v", tc.Matrix))
	}
}
