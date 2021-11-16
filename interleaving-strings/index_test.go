package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_IsInterleave(t *testing.T) {
	testCases := []struct {
		S1       string
		S2       string
		S3       string
		Expected bool
	}{
		{
			"aabcc",
			"dbbca",
			"aadbbcbcac",
			true,
		},
		{
			"aabcc",
			"dbbca",
			"aadbbbaccc",
			false,
		},
		{
			"a",
			"",
			"c",
			false,
		},
		{
			"db",
			"b",
			"cbb",
			false,
		},
		{
			"",
			"",
			"",
			true,
		},
	}
	for _, tc := range testCases {
		actual := IsInterleave(tc.S1, tc.S2, tc.S3)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for Cost %s, %s and %s", tc.S1, tc.S2, tc.S3))
	}
}
