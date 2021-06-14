package algorithms

import (
	"fmt"
	"testing"

	. "github.com/ashinzekene/algorithms/utils/lists"
	"github.com/stretchr/testify/assert"
)

func Test_AddTwoNumbers(t *testing.T) {
	testCases := []struct {
		L1       *ListNode
		L2       *ListNode
		Expected *ListNode
	}{
		{
			ListNodeFromList([]int{1, 5, 6}),
			ListNodeFromList([]int{6, 5, 1}),
			ListNodeFromList([]int{7, 0, 8}),
		},
		{
			ListNodeFromList([]int{2, 4, 3}),
			ListNodeFromList([]int{5, 6, 4}),
			ListNodeFromList([]int{7, 0, 8}),
		},
		{
			ListNodeFromList([]int{6, 7}),
			ListNodeFromList([]int{1, 5}),
			ListNodeFromList([]int{7, 2, 1}),
		},
		{
			ListNodeFromList([]int{5, 5, 5, 5, 5}),
			ListNodeFromList([]int{6, 4, 4, 4}),
			ListNodeFromList([]int{1, 0, 0, 0, 6}),
		},
		{
			ListNodeFromList([]int{9, 9, 9}),
			ListNodeFromList([]int{1, 1, 1, 1, 1, 1, 1, 1, 1}),
			ListNodeFromList([]int{0, 1, 1, 2, 1, 1, 1, 1, 1}),
		},
	}
	for _, tc := range testCases {
		actual := AddTwoNumbers(tc.L1, tc.L2)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v and %v", tc.L1, tc.L2))
	}
}
