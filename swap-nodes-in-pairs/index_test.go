package algorithms

import (
	"testing"

	. "github.com/ashinzekene/algorithms/utils/lists"
	"github.com/stretchr/testify/assert"
)

func Test_3SumClosest(t *testing.T) {
	testCases := []struct {
		ListNode *ListNode
		Expected *ListNode
	}{
		{
			ListNodeFromList([]int{-1, 0, 1, 2, -1, -4}),
			ListNodeFromList([]int{0, -1, 2, 1, -4, -1}),
		},
		{
			ListNodeFromList([]int{1, 2, 3, 4}),
			ListNodeFromList([]int{2, 1, 4, 3}),
		},
		{
			ListNodeFromList([]int{6, -2, 4, 7, 5}),
			ListNodeFromList([]int{-2, 6, 7, 4, 5}),
		},
		{
			ListNodeFromList([]int{1, -1, -1, 0, 2, -3, 2, 4}),
			ListNodeFromList([]int{-1, 1, 0, -1, -3, 2, 4, 2}),
		},
	}
	for _, tc := range testCases {
		actual := swapPairs(tc.ListNode)
		assert.Equal(t, tc.Expected, actual)
	}
}
