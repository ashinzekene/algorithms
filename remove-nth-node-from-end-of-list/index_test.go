package algorithms

import (
	"fmt"
	"testing"

	"github.com/ashinzekene/algorithms/utils/lists"
	"github.com/stretchr/testify/assert"
)

func Test_removeNthFromEnd(t *testing.T) {
	testCases := []struct {
		Node     *lists.ListNode
		Target   int
		Expected *lists.ListNode
	}{
		{
			lists.ListNodeFromString("1,2,3,4,5", ","),
			2,
			lists.ListNodeFromString("1,2,3,5", ","),
		},
		{
			lists.ListNodeFromString("1,2,6,3", ","),
			1,
			lists.ListNodeFromString("1,2,6", ","),
		},
		{
			lists.ListNodeFromString("1,2,6,3,5,7,9,10,12", ","),
			7,
			lists.ListNodeFromString("1,2,3,5,7,9,10,12", ","),
		},
		{
			lists.ListNodeFromString("6,5,3,2,6,3", ","),
			6,
			lists.ListNodeFromString("5,3,2,6,3", ","),
		},
		{
			lists.ListNodeFromString("1", ","),
			1,
			lists.ListNodeFromString("", ","),
		},
	}
	for _, tc := range testCases {
		actual := removeNthFromEnd(tc.Node, tc.Target)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Node))
	}
}
