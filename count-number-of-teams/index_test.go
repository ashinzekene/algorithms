package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumTeams(t *testing.T) {
	testCases := []struct {
		Rating   []int
		Expected int
	}{
		{
			[]int{2, 5, 3, 4, 1},
			3,
		},
		{
			[]int{2, 1, 3},
			0,
		},
		{
			[]int{1, 2, 3, 4},
			4,
		},
		{
			[]int{4, 3, 2, 1},
			4,
		},
		{
			[]int{1},
			0,
		},
		{
			[]int{1, 2},
			0,
		},
		{
			[]int{1, 6, 4, 5},
			1,
		},
	}
	for _, tc := range testCases {
		result := NumTeams(tc.Rating)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for %v,", tc.Rating))
	}
}
