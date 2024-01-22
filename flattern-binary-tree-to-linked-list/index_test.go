package algorithms

import (
	"fmt"
	"testing"

	. "github.com/ashinzekene/algorithms/utils/trees"
	"github.com/stretchr/testify/assert"
)

func Test_buildTree(t *testing.T) {
	testCases := []struct {
		Root     *TreeNode
		Expected *TreeNode
	}{
		{
			TreeFromListLevelOrder([]int{1, 2, 5, 3, 4, -101, 6}, -101).Head,
			TreeFromListLevelOrder([]int{1, 2, 5, 3, 4, -101, 6}, -101).Head,
		},
	}
	for _, tc := range testCases {
		flatten(tc.Root)
		fmt.Println((&Tree{Head: tc.Root}).LevelOrder())
		assert.Equal(t, tc.Expected, tc.Root, fmt.Sprintf("Failed for %v", tc.Root))
	}
}
