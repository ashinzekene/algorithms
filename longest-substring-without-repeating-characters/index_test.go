package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LengthOfLongestSubstring(t *testing.T) {
	testCases := []struct {
		S        string
		Expected int
	}{
		{
			"abcabcbb",
			3,
		},
		{
			"pwwkew",
			3,
		},
		{
			"dvdf",
			3,
		},
		{
			"bbbbbbbbb",
			1,
		},
		{
			"",
			0,
		},
	}
	for _, tc := range testCases {
		actual := LengthOfLongestSubstring(tc.S)
		assert.Equal(t, tc.Expected, actual)
	}
}
