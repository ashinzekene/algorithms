package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LetterCasePermutation(t *testing.T) {
	testCases := []struct {
		S        string
		Expected []string
	}{
		{
			"a1B1",
			[]string{"a1b1", "A1b1", "a1B1", "A1B1"},
		},
		{
			"234567",
			[]string{"234567"},
		},
		{
			"a111",
			[]string{"a111", "A111"},
		},
		{
			"11B1",
			[]string{"11B1", "11b1"},
		},
		{
			"bbb",
			[]string{"bbb", "Bbb", "bBb", "bbB", "BbB", "BBB", "bBB", "BBb"},
		},
	}
	for _, tc := range testCases {
		actual := LetterCasePermutation(tc.S)
		assert.ElementsMatch(t, tc.Expected, actual, "Did not work for "+tc.S)
	}
}
