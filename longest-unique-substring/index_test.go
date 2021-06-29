package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLongestSubstring1(t *testing.T) {
	testCases := []struct {
		S        string
		Expected int
	}{
		{
			"abcabcbb",
			3,
		},
		{
			"abcabcdabcdef",
			6,
		},
		{
			"pwwkew",
			3,
		},
		{
			"bbbbb",
			1,
		},
		{
			"",
			0,
		},
		{
			"aab",
			2,
		},
		{
			"ab",
			2,
		},
		{
			"dvdf",
			3,
		},
		{
			"cdd",
			2,
		},
		{
			"anviaj",
			5,
		},
	}
	for _, tc := range testCases {
		actual := lengthOfLongestSubstring1(tc.S)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("lengthOfLongestSubstring1 did not work for %v", tc.S))
		actual = lengthOfLongestSubstring(tc.S)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("lengthOfLongestSubstring did not work for %v", tc.S))
	}
}
