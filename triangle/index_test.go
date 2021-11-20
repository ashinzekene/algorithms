package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MinimumTotal(t *testing.T) {
	testCases := []struct {
		Triangle [][]int
		Expected int
	}{
		{
			[][]int{
				{2},
				{3, 4},
				{6, 5, 7},
				{4, 1, 8, 3},
			},
			11,
		},
		{
			[][]int{
				{2},
				{1, 4},
				{6, -5, 7},
				{4, 1, 8, 3},
			},
			-1,
		},
		{
			[][]int{
				{-10},
			},
			-10,
		},
		{
			[][]int{
				{-10},
				{11, 10},
			},
			0,
		},
	}
	for _, tc := range testCases {
		actual := MinimumTotal(tc.Triangle)
		assert.Equal(t, tc.Expected, actual)
	}
}
