package algorithms

import (
	"fmt"
	"testing"

	"github.com/ashinzekene/algorithms/utils/lists"
	"github.com/stretchr/testify/assert"
)

func Test_reverseBetween(t *testing.T) {
	testCases := []struct {
		Head     *lists.ListNode
		Start    int
		End      int
		Expected *lists.ListNode
	}{
		{
			lists.ListNodeFromList([]int{1, 2, 3, 4, 5}),
			2,
			4,
			lists.ListNodeFromList([]int{1, 4, 3, 2, 5}),
		},
		{
			lists.ListNodeFromList([]int{1, 2, 3, 4, 5}),
			1,
			4,
			lists.ListNodeFromList([]int{4, 3, 2, 1, 5}),
		},
		{
			lists.ListNodeFromList([]int{1, 2, 3, 4, 5}),
			1,
			5,
			lists.ListNodeFromList([]int{5, 4, 3, 2, 1}),
		},
	}
	for _, tc := range testCases {
		actual := reverseBetween(tc.Head, tc.Start, tc.End)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v, Start: %d, End: %d", tc.Head, tc.Start, tc.End))
	}
}
