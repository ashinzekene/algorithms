package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Permutations(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected [][]int
	}{
		{
			[]int{1, 2},
			[][]int{
				{1, 2},
				{2, 1},
			},
		},
		{
			[]int{1, 2, 3},
			[][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			[]int{1, 1, 2},
			[][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			},
		},
		{
			[]int{1, 2, 1},
			[][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			},
		},
		{
			[]int{1, 2, 3, 4},
			[][]int{
				{1, 2, 3, 4},
				{1, 2, 4, 3},
				{1, 3, 2, 4},
				{1, 3, 4, 2},
				{1, 4, 2, 3},
				{1, 4, 3, 2},
				{2, 1, 3, 4},
				{2, 1, 4, 3},
				{2, 3, 1, 4},
				{2, 3, 4, 1},
				{2, 4, 1, 3},
				{2, 4, 3, 1},
				{3, 1, 2, 4},
				{3, 1, 4, 2},
				{3, 2, 1, 4},
				{3, 2, 4, 1},
				{3, 4, 1, 2},
				{3, 4, 2, 1},
				{4, 1, 2, 3},
				{4, 1, 3, 2},
				{4, 2, 1, 3},
				{4, 2, 3, 1},
				{4, 3, 1, 2},
				{4, 3, 2, 1},
			},
		},
	}
	for i, tc := range testCases {
		result := permuteUnique(tc.Nums)
		assert.ElementsMatch(t, tc.Expected, result, fmt.Sprintf("permute failed for %v", i))
	}
}
