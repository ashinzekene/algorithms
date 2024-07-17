package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_possibleBipartition(t *testing.T) {
	testCases := []struct {
		N        int
		Dislikes [][]int
		Expected bool
	}{
		{
			4,
			[][]int{
				{1, 2}, {1, 3}, {2, 4},
			},
			true,
		},
		{
			3,
			[][]int{
				{1, 2}, {1, 3}, {2, 3},
			},
			false,
		},
		{
			5,
			[][]int{
				{1, 2}, {3, 4}, {4, 5}, {3, 5},
			},
			false,
		},
	}
	for _, tc := range testCases {
		result := possibleBipartition(tc.N, tc.Dislikes)
		assert.ElementsMatch(t, tc.Expected, result, fmt.Sprintf("Failed for %v", tc.Dislikes))
	}
}
