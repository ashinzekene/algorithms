package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CountSquares(t *testing.T) {
	testCases := []struct {
		Matrix   [][]int
		Expected int
	}{
		{
			[][]int{
				{0, 1, 1, 1},
				{1, 1, 1, 1},
				{0, 1, 1, 1},
			},
			15,
		},
		{
			[][]int{
				{1, 0, 1},
				{1, 1, 0},
				{1, 1, 0},
			},
			7,
		},
		{
			[][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 1, 0},
			},
			9,
		},
		{
			[][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 1, 0},
				{1, 1, 1},
				{1, 1, 0},
			},
			8,
		},
		{
			[][]int{
				{1, 1, 1},
				{1, 1, 1},
			},
			8,
		},
		{
			[][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			14,
		},
	}
	for _, tc := range testCases {
		result := CountSquares(tc.Matrix)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for matrix %v", tc.Matrix))
	}
}
