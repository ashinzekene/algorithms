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
				{3, 5},
				{8},
			},
		},
		{
			[]int{10, 1, 2, 7, 6, 1, 5},
			8,
			[][]int{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
		{
			[]int{2, 5, 2, 1, 2},
			5,
			[][]int{
				{1, 2, 2},
				{5},
			},
		},
		{
			[]int{2, 5, 2, 1, 2},
			7,
			[][]int{
				{2, 5},
				{1, 2, 2, 2},
			},
		},
	}
	for _, tc := range testCases {
		result := CombinationSum2(tc.Combinations, tc.Target)
		assert.ElementsMatch(t, tc.Expected, result, fmt.Sprintf("Failed for %v for sum %v", tc.Combinations, tc.Target))
	}
}
