package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PalindromeNumber(t *testing.T) {
	testCases := []struct {
		Number   int
		Expected [][]int
	}{
		{
			3,
			[][]int{
				[]int{1},
				[]int{1, 1},
				[]int{1, 2, 1},
			},
		},
		{
			4,
			[][]int{
				[]int{1},
				[]int{1, 1},
				[]int{1, 2, 1},
				[]int{1, 3, 3, 1},
			},
		},
		{
			1,
			[][]int{
				[]int{1},
			},
		},
		{
			2,
			[][]int{
				[]int{1},
				[]int{1, 1},
			},
		},
		{
			8,
			[][]int{
				[]int{1},
				[]int{1, 1},
				[]int{1, 2, 1},
				[]int{1, 3, 3, 1},
				[]int{1, 4, 6, 4, 1},
				[]int{1, 5, 10, 10, 5, 1},
				[]int{1, 6, 15, 20, 15, 6, 1},
				[]int{1, 7, 21, 35, 35, 21, 7, 1},
			},
		},
		{
			5,
			[][]int{
				[]int{1},
				[]int{1, 1},
				[]int{1, 2, 1},
				[]int{1, 3, 3, 1},
				[]int{1, 4, 6, 4, 1},
			},
		},
	}
	for _, tc := range testCases {
		actual := generate(tc.Number)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Number))
	}
}
