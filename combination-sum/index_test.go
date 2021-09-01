package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum(t *testing.T) {
	testCases := []struct {
		Combinations []int
		Target       int
		Expected     [][]int
	}{
		{
			[]int{2, 3, 5, 8},
			8,
			[][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
				{8},
			},
		},
		{
			[]int{2, 3, 5},
			8,
			[][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
		{
			[]int{2, 3, 5},
			7,
			[][]int{
				{2, 2, 3},
				{2, 5},
			},
		},
		{
			[]int{2, 4, 6, 8},
			8,
			[][]int{
				{2, 2, 2, 2},
				{2, 2, 4},
				{2, 6},
				{4, 4},
				{8},
			},
		},
	}
	for _, tc := range testCases {
		result := combinationSum(tc.Combinations, tc.Target)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for %v for sum %v", tc.Combinations, tc.Target))
	}
}
