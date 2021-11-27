package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subsetsWithDup(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected [][]int
	}{
		{
			[]int{1, 2, 3},
			[][]int{
				{},
				{1},
				{2},
				{1, 2},
				{3},
				{1, 3},
				{2, 3},
				{1, 2, 3},
			},
		},
		{
			[]int{0},
			[][]int{
				{},
				{0},
			},
		},
		{
			[]int{1, 2},
			[][]int{
				{},
				{1},
				{2},
				{1, 2},
			},
		},
		{
			[]int{4, 4, 1, 4},
			[][]int{
				{},
				{1},
				{4},
				{4, 4},
				{1, 4},
				{1, 4, 4},
				{4, 4, 4},
				{1, 4, 4, 4},
			},
		},
		{
			[]int{1, 2, 2},
			[][]int{
				{},
				{1},
				{2},
				{1, 2},
				{2, 2},
				{1, 2, 2},
			},
		},
		{
			[]int{1, 2, 2, 3},
			[][]int{
				{},
				{1},
				{2},
				{3},
				{1, 2},
				{2, 2},
				{2, 3},
				{1, 3},
				{1, 2, 2},
				{2, 2, 3},
				{1, 2, 3},
				{1, 2, 2, 3},
			},
		},
	}
	for _, tc := range testCases {
		result := subsetsWithDup(tc.Nums)
		assert.ElementsMatch(t, tc.Expected, result, fmt.Sprintf("Failed for %v", tc.Nums))
	}
}
