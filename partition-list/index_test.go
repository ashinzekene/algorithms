package algorithms

import (
	"fmt"
	"testing"

	. "github.com/ashinzekene/algorithms/utils/lists"
	"github.com/stretchr/testify/assert"
)

func Test_Partition(t *testing.T) {
	testCases := []struct {
		Head     *ListNode
		Pivot    int
		Expected *ListNode
	}{
		{
			ListNodeFromList([]int{1, 4, 3, 2, 5, 2}),
			3,
			ListNodeFromList([]int{1, 2, 2, 4, 3, 5}),
		},
		{
			ListNodeFromList([]int{2, 1}),
			2,
			ListNodeFromList([]int{1, 2}),
		},
		{
			ListNodeFromList([]int{1, 2, 3, 4, 5, 6, 7}),
			8,
			ListNodeFromList([]int{1, 2, 3, 4, 5, 6, 7}),
		},
		{
			ListNodeFromList([]int{1, 2, 3, 4, 5, 6, 7}),
			5,
			ListNodeFromList([]int{1, 2, 3, 4, 5, 6, 7}),
		},
		{
			ListNodeFromList([]int{1, 2, 3, 4, 5, 6, 7}),
			1,
			ListNodeFromList([]int{1, 2, 3, 4, 5, 6, 7}),
		},
		{
			ListNodeFromList([]int{}),
			1,
			ListNodeFromList([]int{}),
		},
		{
			ListNodeFromList([]int{7, 6, 5, 4, 3, 2, 2, 1}),
			7,
			ListNodeFromList([]int{6, 5, 4, 3, 2, 2, 1, 7}),
		},
	}
	for _, tc := range testCases {
		actual := Partition(tc.Head, tc.Pivot)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Head))
	}
}
