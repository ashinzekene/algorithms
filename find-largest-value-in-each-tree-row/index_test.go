package algorithms

import (
	"fmt"
	"testing"

	. "github.com/ashinzekene/algorithms/utils/trees"
	"github.com/stretchr/testify/assert"
)

func Test_FindDigits(t *testing.T) {
	testCases := []struct {
		Root    *TreeNode
		Expected []int
	}{
		{
			TreeFromListLevelOrder([]int{1,3,2,5,3,-1,9}, -1).Head,
			[]int{1,3, 9},
		},
		{
			TreeFromListLevelOrder([]int{1,3}, -1).Head,
			[]int{1,3},
		},
		{
			TreeFromListLevelOrder([]int{1}, -1).Head,
			[]int{1},
		},
		{
			TreeFromListLevelOrder([]int{}, -1).Head,
			[]int{},
		},
		{
			TreeFromListLevelOrder([]int{-1, 10, -20}, -1).Head,
			[]int{-1, -20},
		},
	}
	for _, tc := range testCases {
		actual := largestValues(tc.Root)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for Digit %v", tc.Root))
	}
}
