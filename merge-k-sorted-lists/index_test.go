package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeKLists(t *testing.T) {
	testCases := []struct {
		Lists    []*ListNode
		Expected *ListNode
	}{
		{
			[]*ListNode{
				ListNodeFromString("1,2,3,4,5", ","),
				ListNodeFromString("1,2,3,5", ","),
			},
			ListNodeFromString("1,1,2,2,3,3,4,5,5", ","),
		},
		{
			[]*ListNode{
				ListNodeFromString("1,2,3,6", ","),
				ListNodeFromString("2,4,5", ","),
				ListNodeFromString("2,5,8", ","),
			},
			ListNodeFromString("1,2,2,2,3,4,5,5,6,8", ","),
		},
		{
			[]*ListNode{
				ListNodeFromString("1,2,6,3,5,7,9,10,12", ","),
				ListNodeFromString("", ","),
				ListNodeFromString("", ","),
			},
			ListNodeFromString("1,2,6,3,5,7,9,10,12", ","),
		},
		{
			[]*ListNode{
				ListNodeFromString("", ","),
				ListNodeFromString("3", ","),
				ListNodeFromString("2", ","),
				ListNodeFromString("1", ","),
			},
			ListNodeFromString("1,2,3", ","),
		},
		{
			[]*ListNode{
				ListNodeFromString("1,2,4", ","),
				ListNodeFromString("1,3,6", ","),
				ListNodeFromString("2,5,8", ","),
				ListNodeFromString("1", ","),
				ListNodeFromString("", ","),
			},
			ListNodeFromString("1,1,1,2,2,3,4,5,6,8", ","),
		},
	}
	for _, tc := range testCases {
		actual := mergeKLists(tc.Lists)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Lists))
	}
}
