package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LeftRotation(t *testing.T) {
	testCases := []struct {
		Nums     [][]int
		Expected [][]int
	}{
		{
			[][]int{
				{0, 1, 2, 0},
				{3, 4, 5, 2},
				{1, 3, 1, 5},
			},
			[][]int{
				{0, 0, 0, 0},
				{0, 4, 5, 0},
				{0, 3, 1, 0},
			},
		},
		{
			[][]int{
				{1, 3},
				{0, 6},
			},
			[][]int{
				{0, 3},
				{0, 0},
			},
		},
		{
			[][]int{
				{1, 4},
				{2, 3},
			},
			[][]int{
				{1, 4},
				{2, 3},
			},
		},
		{
			[][]int{
				{1},
			},
			[][]int{
				{1},
			},
		},
		{
			[][]int{
				{1, 4},
				{0, 0},
			},
			[][]int{
				{0, 0},
				{0, 0},
			},
		},
	}
	for _, tc := range testCases {
		actual := setZeroes(tc.Nums)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Nums))
	}
}
