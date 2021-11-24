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
			"aacaba",
			2,
		},
		{
			"abcd",
			1,
		},
		{
			"aaaaa",
			4,
		},
		{
			"acbadbaada",
			2,
		},
	}
	for _, tc := range testCases {
		actual := NumSplits(tc.S)
		assert.Equal(t, tc.Expected, actual, "Did not work for "+tc.S)
	}
}
