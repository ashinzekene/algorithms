package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PermutationEquation(t *testing.T) {
	testCases := []struct {
		Array    []int
		Expected []int
	}{
		{
			[]int{2, 3, 2, 4},
			[]int{2, 3, 2, 5},
		},
		{
			[]int{2, 3, 4},
			[]int{2, 3, 5},
		},
		{
			[]int{2, 0, 1},
			[]int{2, 0, 2},
		},
		{
			[]int{1, 9},
			[]int{2, 0},
		},
		{
			[]int{9, 9},
			[]int{1, 0, 0},
		},
		{
			[]int{9},
			[]int{1, 0},
		},
		{
			[]int{0},
			[]int{1},
		},
	}
	for _, tc := range testCases {
		actual := plusOne(tc.Array)
		assert.Equal(t, tc.Expected, actual)
	}
}
