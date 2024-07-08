package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMinHeightTrees(t *testing.T) {
	testCases := []struct {
		N int
		Edges    [][]int
		Expected []int
	}{
		{
			4,
			[][]int{
			{1,0},
			{1,2},
			{1,3},
			},
			[]int{1},
		},
		{
			6,
			[][]int{
			{3,0},
			{3,1},
			{3,2},
			{3,4},
			{5,4},
			},
			[]int{3,4},
		},
		{
			3,
			[][]int{
			{1,0},
			{1,2},
			},
			[]int{1},
		},
		{
			11,
			[][]int{
			{0,1},
			{0,2},
			{2,3},
			{0,4},
			{2,5},
			{5,6},
			{3,7},
			{6,8},
			{8,9},
			{9,10},
			},
			[]int{5,6},
		},
		{
		},
	}
	for _, tc := range testCases {
		actual := findMinHeightTrees(tc.N, tc.Edges)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Edges))
	}
}
