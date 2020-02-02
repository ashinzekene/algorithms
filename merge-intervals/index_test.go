package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LeftRotation(t *testing.T) {
	testCases := []struct {
		Nums     [][]int
		Expected [][]int
	}{
		{
			[][]int{
				[]int{2, 6},
				[]int{8, 10},
				[]int{15, 18},
			},
			[][]int{
				[]int{2, 6},
				[]int{8, 10},
				[]int{15, 18},
			},
		},
		{
			[][]int{
				[]int{1, 3},
				[]int{2, 6},
				[]int{6, 10},
				[]int{15, 18},
			},
			[][]int{
				[]int{1, 10},
				[]int{15, 18},
			},
		},
		{
			[][]int{
				[]int{1, 4},
				[]int{2, 3},
			},
			[][]int{
				[]int{1, 4},
			},
		},
		{
			[][]int{
				[]int{1, 4},
				[]int{0, 4},
			},
			[][]int{
				[]int{0, 4},
			},
		},
		{
			[][]int{
				[]int{1, 4},
				[]int{0, 0},
			},
			[][]int{
				[]int{0, 0},
				[]int{1, 4},
			},
		},
		{
			[][]int{
				[]int{2, 3},
				[]int{4, 5},
				[]int{6, 7},
				[]int{8, 9},
				[]int{1, 10},
			},
			[][]int{
				[]int{1, 10},
			},
		},
	}
	for _, tc := range testCases {
		actual := merge(tc.Nums)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Nums))
	}
}
