package algorithms

import (
	"fmt"
	"testing"

	. "github.com/ashinzekene/algorithms/utils/trees"
	"github.com/stretchr/testify/assert"
)

func Test_generateTrees(t *testing.T) {
	testCases := []struct {
		N        int
		Expected []*TreeNode
	}{
		{
			2,
			[]*TreeNode{
				TreeFromListLevelOrder([]int{1, -1001, 2}, -1001).Head,
				TreeFromListLevelOrder([]int{2, 1, -1001}, -1001).Head,
			},
		},
		{
			1,
			[]*TreeNode{
				TreeFromListLevelOrder([]int{1}, -1001).Head,
			},
		},
	}
	for _, tc := range testCases {
		actual := generateTrees(tc.N)
		assert.ElementsMatch(t, tc.Expected, actual, "Did not work for "+fmt.Sprint(tc.N))
	}
}
