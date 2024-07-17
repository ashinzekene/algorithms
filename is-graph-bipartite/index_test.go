package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isBipartite(t *testing.T) {
	testCases := []struct {
		Graph [][]int
		Expected bool
	}{
		{
			[][]int{
				{1,2,3},{0,2},{0,1,3},{0,2},
			},
			false,
		},
		{
			[][]int{
				{1,3},{0,2},{1,3},{0,2},
			},
			true,
		},
		{
			[][]int{
				{1},{0},{3},{2},
			},
			true,
		},
	}
	for _, tc := range testCases {
		result := isBipartite(tc.Graph)
		assert.ElementsMatch(t, tc.Expected, result, fmt.Sprintf("Failed for %v", tc.Graph))
	}
}
