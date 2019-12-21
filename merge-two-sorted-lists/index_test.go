package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mergeTwoLists(t *testing.T) {
	testCases := []struct {
		Node1    *ListNode
		Node2    *ListNode
		Expected *ListNode
	}{
		{
			ListNodeFromString("1,2,3,4,5", ","),
			ListNodeFromString("1,2,3,5", ","),
			ListNodeFromString("1,1,2,2,3,3,4,5,5", ","),
		},
		{
			ListNodeFromString("1,2,3,6", ","),
			ListNodeFromString("2,4,5", ","),
			ListNodeFromString("1,2,2,3,4,5,6", ","),
		},
		{
			ListNodeFromString("1,2,6,3,5,7,9,10,12", ","),
			ListNodeFromString("", ","),
			ListNodeFromString("1,2,6,3,5,7,9,10,12", ","),
		},
		{
			ListNodeFromString("", ","),
			ListNodeFromString("3", ","),
			ListNodeFromString("3", ","),
		},
		{
			ListNodeFromString("1", ","),
			ListNodeFromString("", ","),
			ListNodeFromString("1", ","),
		},
	}
	for _, tc := range testCases {
		actual := mergeTwoLists(tc.Node1, tc.Node2)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Node1))
	}
}
