package algorithms

import (
	"fmt"
	"testing"

	. "github.com/ashinzekene/algorithms/utils/lists"
	"github.com/stretchr/testify/assert"
)

func Test_swapNodes(t *testing.T) {
	testCases := []struct {
		ListNode *ListNode
		K        int
		Expected *ListNode
	}{
		{
			ListNodeFromList([]int{-1, 0, 1, 2, -1, -4}),
			2,
			ListNodeFromList([]int{-1, -1, 1, 2, 0, -4}),
		},
		{
			ListNodeFromList([]int{1, 2, 3, 4}),
			3,
			ListNodeFromList([]int{1, 3, 2, 4}),
		},
		{
			ListNodeFromList([]int{6, -2, 4, 7, 5}),
			1,
			ListNodeFromList([]int{5, -2, 4, 7, 6}),
		},
		{
			ListNodeFromList([]int{1, -1, -1, 0, 2, -3, 2, 4}),
			5,
			ListNodeFromList([]int{1, -1, -1, 2, 0, -3, 2, 4}),
		},
	}
	for _, tc := range testCases {
		actual := swapNodes(tc.ListNode, tc.K)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("swapPairs did not work for %v", tc.ListNode))
	}
}
