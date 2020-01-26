package algorithms

import (
	"fmt"
	"testing"

	"github.com/ashinzekene/algorithms/utils/trees"
	"github.com/stretchr/testify/assert"
)

func Test_LeftRotation(t *testing.T) {
	testCases := []struct {
		Nums     *trees.TreeNode
		Expected int
	}{
		{
			trees.ArrayToTree([]int{3, 2, 6, 4}).Head,
			3,
		},
		{
			trees.ArrayToTree([]int{9, 3, 20, 15, 7}).Head,
			3,
		},
	}
	for _, tc := range testCases {
		actual := maxDepth(tc.Nums)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Nums))
	}
}
