package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CutTheSticks(t *testing.T) {
	testCases := []struct {
		String   string
		Expected string
	}{
		{
			"dnotq",
			"crime",
		},
		{
			"flgxswdliefy",
			"encyclopedia",
		},
		{
			"rajsb",
			"qqqqq",
		},
		{
			"bvqmjhgghjmqvbiqzjugthwmdv",
			"abcdefghijklmnopqrstuvwxyz",
		},
		{
			"eobamwpnlmhklrq",
			"drugtrafficking",
		},
	}
	for _, tc := range testCases {
		result := DecryptMessage(tc.String)
		assert.Equal(t, tc.Expected, result, "Failed for %v", tc.String)
	}
}
