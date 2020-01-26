package algorithms

import (
	"fmt"
	"testing"

	. "github.com/ashinzekene/algorithms/utils/trees"
	"github.com/stretchr/testify/assert"
)

func Test_Squares(t *testing.T) {
	testCases := []struct {
		Tree     *TreeNode
		Expected bool
	}{
		{
			&TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
					Right: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
					Left: &TreeNode{
						Val: 4,
					},
				},
			},
			true,
		},
		{
			ArrayToTree([]int{1, 2, 2, 3, 4, 4, 3}).Head,
			false,
		},
		{
			ArrayToTree([]int{1, 2, 9, 5, 6, 7, 3}).Head,
			false,
		},
	}
	for _, tc := range testCases {
		result := isSymmetric(tc.Tree)
		result2 := isSymmetric2(tc.Tree)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for %#v", tc.Tree))
		assert.Equal(t, tc.Expected, result2, fmt.Sprintf("Failed for %#v", tc.Tree))
	}
}
