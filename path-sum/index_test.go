package algorithms

import (
	"fmt"
	"testing"

	. "github.com/ashinzekene/algorithms/utils/trees"
	"github.com/stretchr/testify/assert"
)

func Test_HasPathSum(t *testing.T) {
	testCases := []struct {
		Root     *Tree
		Sum      int
		Expected bool
	}{
		{
			TreeFromListLevelOrder([]int{5, 4, 8, 11, -1001, 13, 4, 7, 2, -1001, -1001, -1001, -1001, 5, 1}, -1001),
			22,
			true,
		},
		{
			TreeFromListLevelOrder([]int{1, 2}, -1001),
			1,
			false,
		},
		{
			TreeFromListLevelOrder([]int{-2, -1001, -3}, -1001),
			-5,
			true,
		},
		{
			TreeFromListLevelOrder([]int{1, 2, 3}, -1001),
			5,
			false,
		},
	}
	for _, tc := range testCases {
		actual := HasPathSum(tc.Root.Head, tc.Sum)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Root.LevelOrder()))
	}
}
