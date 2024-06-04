package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum(t *testing.T) {
	testCases := []struct {
		Combinations int
		Target       int
		Expected     [][]int
	}{
		{
			3,
			7,
			[][]int{
				{1,2,4},
			},
		},
		{
			3,
			9,
			[][]int{
				{1,2,6},
				{1,3,5},
				{2,3,4},
			},
		},
		{
			4,
			24,
			[][]int{
				{1,6,8,9},
				{2,5,8,9},
				{2,6,7,9},
				{3,4,8,9},
				{3,5,7,9},
				{3,6,7,8},
				{4,5,6,9},
				{4,5,7,8},
			},
		},
	}
	for _, tc := range testCases {
		result := combinationSum3(tc.Combinations, tc.Target)
		assert.ElementsMatch(t, tc.Expected, result, fmt.Sprintf("Failed for %v for sum %v", tc.Combinations, tc.Target))
	}
}
