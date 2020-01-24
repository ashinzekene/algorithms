package algorithms

import (
	"fmt"
	"testing"

	"github.com/ashinzekene/algorithms/utils/lists"
	"github.com/stretchr/testify/assert"
)

func Test_mergeTwoLists(t *testing.T) {
	testCases := []struct {
		Node1    *lists.ListNode
		Node2    *lists.ListNode
		Expected *lists.ListNode
	}{
		{
			lists.ListNodeFromString("1,2,3,4,5", ","),
			lists.ListNodeFromString("1,2,3,5", ","),
			lists.ListNodeFromString("1,1,2,2,3,3,4,5,5", ","),
		},
		{
			lists.ListNodeFromString("1,2,3,6", ","),
			lists.ListNodeFromString("2,4,5", ","),
			lists.ListNodeFromString("1,2,2,3,4,5,6", ","),
		},
		{
			lists.ListNodeFromString("1,2,6,3,5,7,9,10,12", ","),
			lists.ListNodeFromString("", ","),
			lists.ListNodeFromString("1,2,6,3,5,7,9,10,12", ","),
		},
		{
			lists.ListNodeFromString("", ","),
			lists.ListNodeFromString("3", ","),
			lists.ListNodeFromString("3", ","),
		},
		{
			lists.ListNodeFromString("1", ","),
			lists.ListNodeFromString("", ","),
			lists.ListNodeFromString("1", ","),
		},
	}
	for _, tc := range testCases {
		actual := mergeTwoLists(tc.Node1, tc.Node2)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Node1))
	}
}
