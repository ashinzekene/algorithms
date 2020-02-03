package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NextPermutation(t *testing.T) {
	testCases := []struct {
		Array    []int
		Expected []int
	}{
		{
			[]int{2, 3, 1, 3, 3},
			[]int{2, 3, 3, 1, 3},
		},
		{
			[]int{1, 5, 1},
			[]int{5, 1, 1},
		},
		{
			[]int{1, 2, 3},
			[]int{1, 3, 2},
		},
		{
			[]int{3, 2, 1},
			[]int{1, 2, 3},
		},
		{
			[]int{3, 2, 3},
			[]int{3, 3, 2},
		},
		{
			[]int{1, 1, 5},
			[]int{1, 5, 1},
		},
		{
			[]int{1, 3, 2},
			[]int{2, 1, 3},
		},
		{
			[]int{1, 3, 5, 4, 2},
			[]int{1, 4, 2, 3, 5},
		},
		{
			[]int{1, 3, 8, 6, 5, 5, 4, 2},
			[]int{1, 4, 2, 3, 5, 5, 6, 8},
		},
	}
	for _, tc := range testCases {
		actual := nextPermutation(tc.Array)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Array))
	}
}
