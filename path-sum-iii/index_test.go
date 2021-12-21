package algorithms

import (
	"fmt"
	"testing"

	. "github.com/ashinzekene/algorithms/utils/trees"
	"github.com/stretchr/testify/assert"
)

func Test_pathSum(t *testing.T) {
	testCases := []struct {
		Root     *Tree
		Sum      int
		Expected int
	}{
		{
			TreeFromListLevelOrder([]int{5, 4, 8, 11, -1001, 13, 4, 7, 2, -1001, -1001, -1001, -1001, 5, 1}, -1001),
			22,
			3,
		},
		{
			TreeFromListLevelOrder([]int{10, 5, -3, 3, 2, -1001, 11, 3, -2, -1001, 1}, -1001),
			8,
			3,
		},
		{
			TreeFromListLevelOrder([]int{1, 2}, -1001),
			1,
			1,
		},
		{
			TreeFromListLevelOrder([]int{-2, -1001, -3}, -1001),
			-5,
			1,
		},
		{
			TreeFromListLevelOrder([]int{0, 1, 1}, -1001),
			1,
			4,
		},
		{
			TreeFromListLevelOrder([]int{1}, -1001),
			0,
			0,
		},
	}
	for _, tc := range testCases {
		actual := pathSum(tc.Root.Head, tc.Sum)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Root.LevelOrder()))
	}
}
