package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeNthFromEnd(t *testing.T) {
	testCases := []struct {
		Node     *ListNode
		Target   int
		Expected *ListNode
	}{
		{
			ListNodeFromString("1,2,3,4,5", ","),
			2,
			ListNodeFromString("1,2,3,5", ","),
		},
		{
			ListNodeFromString("1,2,6,3", ","),
			1,
			ListNodeFromString("1,2,6", ","),
		},
		{
			ListNodeFromString("1,2,6,3,5,7,9,10,12", ","),
			7,
			ListNodeFromString("1,2,3,5,7,9,10,12", ","),
		},
		{
			ListNodeFromString("6,5,3,2,6,3", ","),
			6,
			ListNodeFromString("5,3,2,6,3", ","),
		},
		{
			ListNodeFromString("1", ","),
			1,
			ListNodeFromString("", ","),
		},
	}
	for _, tc := range testCases {
		actual := removeNthFromEnd(tc.Node, tc.Target)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Node))
	}
}
