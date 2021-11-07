package algorithms

import (
	"fmt"
	"testing"

	. "github.com/ashinzekene/algorithms/utils/lists"
	"github.com/stretchr/testify/assert"
)

func Test_RotateList(t *testing.T) {
	testCases := []struct {
		L1       *ListNode
		K        int
		Expected *ListNode
	}{
		{
			ListNodeFromList([]int{1, 2, 3, 4, 5, 6}),
			2,
			ListNodeFromList([]int{5, 6, 1, 2, 3, 4}),
		},
		{
			ListNodeFromList([]int{2, 4, 3}),
			0,
			ListNodeFromList([]int{2, 4, 3}),
		},
		{
			ListNodeFromList([]int{6, 7}),
			2,
			ListNodeFromList([]int{6, 7}),
		},
		{
			ListNodeFromList([]int{1, 2, 3, 4, 5}),
			10,
			ListNodeFromList([]int{1, 2, 3, 4, 5}),
		},
		{
			ListNodeFromList([]int{}),
			2,
			ListNodeFromList([]int{}),
		},
	}
	for _, tc := range testCases {
		actual := RotateRight(tc.L1, tc.K)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v and %v", tc.L1, tc.K))
	}
}
