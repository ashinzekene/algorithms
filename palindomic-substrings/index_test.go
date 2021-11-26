package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NextPermutation(t *testing.T) {
	testCases := []struct {
		S        string
		Expected int
	}{
		{
			"abc",
			3,
		},
		{
			"abcd",
			4,
		},
		{
			"aaa",
			6,
		},
		{
			"acbadbaada",
			12,
		},
	}
	for _, tc := range testCases {
		actual := CountSubstrings(tc.S)
		assert.Equal(t, tc.Expected, actual, "Did not work for "+tc.S)
	}
}
