package algorithms

import (
	"fmt"
	"testing"

	. "github.com/ashinzekene/algorithms/utils/trees"
	"github.com/stretchr/testify/assert"
)

func Test_buildTree(t *testing.T) {
	testCases := []struct {
		InOrder   []int
		PostOrder []int
		Expected  *TreeNode
	}{
		{
			[]int{9, 3, 15, 20, 7},
			[]int{9, 15, 7, 20, 3},
			TreeFromListLevelOrder([]int{3, 9, 20, -3001, -3001, 15, 7}, -3001).Head,
		},
		{
			[]int{-1},
			[]int{-1},
			TreeFromListLevelOrder([]int{-1}, -3001).Head,
		},
	}
	for _, tc := range testCases {
		result := buildTree(tc.InOrder, tc.PostOrder)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for %v and %v", tc.InOrder, tc.PostOrder))
	}
}
