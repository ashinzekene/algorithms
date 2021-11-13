package algorithms

import (
	"fmt"
	"testing"

	. "github.com/ashinzekene/algorithms/utils/lists"
	"github.com/stretchr/testify/assert"
)

func Test_DeleteDuplicates(t *testing.T) {
	testCases := []struct {
		ListNode *ListNode
		Expected *ListNode
	}{
		{
			ListNodeFromList([]int{1, 2, 3, 3, 4, 4, 5}),
			ListNodeFromList([]int{1, 2, 5}),
		},
		{
			ListNodeFromList([]int{1, 1, 1, 1, 2, 3, 4}),
			ListNodeFromList([]int{2, 3, 4}),
		},
		{
			ListNodeFromList([]int{}),
			ListNodeFromList([]int{}),
		},
		{
			ListNodeFromList([]int{2, 2, 2}),
			ListNodeFromList([]int{}),
		},
		{
			ListNodeFromList([]int{1, 2, 2, 2}),
			ListNodeFromList([]int{1}),
		},
		{
			ListNodeFromList([]int{2, 1, 2}),
			ListNodeFromList([]int{2, 1, 2}),
		},
		{
			ListNodeFromList([]int{2, 2, 1, 2}),
			ListNodeFromList([]int{1, 2}),
		},
	}
	for _, tc := range testCases {
		actual := DeleteDuplicates(tc.ListNode)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Delete duplicate did not work for %v", tc.ListNode))
	}
}
