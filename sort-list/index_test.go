package algorithms

import (
	"fmt"
	"testing"

	. "github.com/ashinzekene/algorithms/utils/lists"
	"github.com/stretchr/testify/assert"
)

func Test_sortList(t *testing.T) {
	testCases := []struct {
		Head     *ListNode
		Expected *ListNode
	}{
		{
			ListNodeFromList([]int{1, 3, 2, 4}),
			ListNodeFromList([]int{1, 2, 3, 4}),
		},
		{
			ListNodeFromList([]int{1, 3}),
			ListNodeFromList([]int{1, 3}),
		},
		{
			ListNodeFromList([]int{}),
			ListNodeFromList([]int{}),
		},
		{
			ListNodeFromList([]int{1}),
			ListNodeFromList([]int{1}),
		},
		{
			ListNodeFromList([]int{1, 4, 7, 5, 3, 6, 8, 4, 3, 2}),
			ListNodeFromList([]int{1, 2, 3, 3, 4, 4, 5, 6, 7, 8}),
		},
	}
	for _, tc := range testCases {
		actual := sortList(tc.Head)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for %s", tc.Head))
	}
}
