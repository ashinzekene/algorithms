package algorithms

import (
	"fmt"
	"testing"

	. "github.com/ashinzekene/algorithms/utils/trees"
	"github.com/stretchr/testify/assert"
)

func Test_PathSum(t *testing.T) {
	testCases := []struct {
		Root     *Tree
		Sum      int
		Expected [][]int
	}{
		{
			TreeFromListLevelOrder([]int{5, 4, 8, 11, -1001, 13, 4, 7, 2, -1001, -1001, -1001, -1001, 5, 1}, -1001),
			22,
			[][]int{
				{5, 4, 11, 2},
				{5, 8, 4, 5},
			},
		},
		{
			TreeFromListLevelOrder([]int{1, 2}, -1001),
			1,
			[][]int{},
		},
		{
			TreeFromListLevelOrder([]int{-2, -1001, -3}, -1001),
			-5,
			[][]int{
				{-2, -3},
			},
		},
		{
			TreeFromListLevelOrder([]int{1, 2, 3}, -1001),
			5,
			[][]int{},
		},
	}
	for _, tc := range testCases {
		actual := PathSum(tc.Root.Head, tc.Sum)
		assert.ElementsMatch(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Root.LevelOrder()))
	}
}
