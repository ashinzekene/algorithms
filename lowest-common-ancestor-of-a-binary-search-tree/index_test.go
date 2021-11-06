package algorithms

import (
	"fmt"
	"testing"

	. "github.com/ashinzekene/algorithms/utils/trees"
	"github.com/stretchr/testify/assert"
)

func Test_LowestCommonAncestor(t *testing.T) {
	testCases := []struct {
		Root     *TreeNode
		P        *TreeNode
		Q        *TreeNode
		Expected *TreeNode
	}{
		{
			TreeFromListLevelOrder([]int{6, 2, 8, 0, 4, 7, 9, -100, 100, 3, 5}, -100).Head,
			TreeFromListLevelOrder([]int{2}, -100).Head,
			TreeFromListLevelOrder([]int{4}, -100).Head,
			TreeFromListLevelOrder([]int{2}, -100).Head,
		},
		{
			TreeFromListLevelOrder([]int{2, 1}, -100).Head,
			TreeFromListLevelOrder([]int{2}, -100).Head,
			TreeFromListLevelOrder([]int{1}, -100).Head,
			TreeFromListLevelOrder([]int{2}, -100).Head,
		},
	}
	for _, tc := range testCases {
		actual := LowestCommonAncestor(tc.Root, tc.P, tc.Q)
		assert.Equal(t, tc.Expected.Val, actual.Val, fmt.Sprintf("Did not work for %v", tc.Root))
	}
}
