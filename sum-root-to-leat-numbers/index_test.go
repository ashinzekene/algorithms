package algorithms

import (
	"fmt"
	"testing"

	. "github.com/ashinzekene/algorithms/utils/trees"
	"github.com/stretchr/testify/assert"
)

func Test_sumNumbers(t *testing.T) {
	testCases := []struct {
		Root     *TreeNode
		Expected int
	}{
		{
			TreeFromListLevelOrder([]int{1, 2, 3}, -1000).Head,
			25,
		},
		{
			TreeFromListLevelOrder([]int{4, 9, 0, 5, 1, -1000, -1000}, -1000).Head,
			1026,
		},
		{
			TreeFromListLevelOrder([]int{4}, -1000).Head,
			4,
		},
		{
			TreeFromListLevelOrder([]int{}, -1000).Head,
			0,
		},
	}
	for _, tc := range testCases {
		result := sumNumbers(tc.Root)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for %v", tc.Root))
	}
}
