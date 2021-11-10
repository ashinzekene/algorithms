package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum(t *testing.T) {
	testCases := []struct {
		N        int
		K        int
		Expected [][]int
	}{
		{
			4,
			2,
			[][]int{
				{2, 4},
				{3, 4},
				{2, 3},
				{1, 2},
				{1, 3},
				{1, 4},
			},
		},
		{
			1,
			1,
			[][]int{
				{1},
			},
		},
		{
			5,
			4,
			[][]int{
				{1, 2, 3, 4},
				{1, 2, 3, 5},
				{1, 2, 4, 5},
				{1, 3, 4, 5},
				{2, 3, 4, 5},
			},
		},
	}
	for _, tc := range testCases {
		result := Combine(tc.N, tc.K)
		assert.ElementsMatch(t, tc.Expected, result, fmt.Sprintf("Failed for %v for sum %v", tc.N, tc.K))
	}
}
