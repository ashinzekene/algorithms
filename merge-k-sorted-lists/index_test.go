package algorithms

import (
	"fmt"
	"testing"

	"github.com/ashinzekene/algorithms/utils/lists"
	"github.com/stretchr/testify/assert"
)

func Test_mergeKLists(t *testing.T) {
	testCases := []struct {
		Lists    []*lists.ListNode
		Expected *lists.ListNode
	}{
		{
			[]*lists.ListNode{
				lists.ListNodeFromString("1,2,3,4,5", ","),
				lists.ListNodeFromString("1,2,3,5", ","),
			},
			lists.ListNodeFromString("1,1,2,2,3,3,4,5,5", ","),
		},
		{
			[]*lists.ListNode{
				lists.ListNodeFromString("1,2,3,6", ","),
				lists.ListNodeFromString("2,4,5", ","),
				lists.ListNodeFromString("2,5,8", ","),
			},
			lists.ListNodeFromString("1,2,2,2,3,4,5,5,6,8", ","),
		},
		{
			[]*lists.ListNode{
				lists.ListNodeFromString("1,2,6,3,5,7,9,10,12", ","),
				lists.ListNodeFromString("", ","),
				lists.ListNodeFromString("", ","),
			},
			lists.ListNodeFromString("1,2,6,3,5,7,9,10,12", ","),
		},
		{
			[]*lists.ListNode{
				lists.ListNodeFromString("", ","),
				lists.ListNodeFromString("3", ","),
				lists.ListNodeFromString("2", ","),
				lists.ListNodeFromString("1", ","),
			},
			lists.ListNodeFromString("1,2,3", ","),
		},
		{
			[]*lists.ListNode{
				lists.ListNodeFromString("1,2,4", ","),
				lists.ListNodeFromString("1,3,6", ","),
				lists.ListNodeFromString("2,5,8", ","),
				lists.ListNodeFromString("1", ","),
				lists.ListNodeFromString("", ","),
			},
			lists.ListNodeFromString("1,1,1,2,2,3,4,5,6,8", ","),
		},
	}
	for _, tc := range testCases {
		actual := mergeKLists(tc.Lists)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Lists))
	}
}
