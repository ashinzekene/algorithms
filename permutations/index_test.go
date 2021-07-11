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
		result := permute(tc.Nums)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("permute failed for %v", i))
		result = permute1(tc.Nums)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("permute1 failed for %v", i))
	}
}
