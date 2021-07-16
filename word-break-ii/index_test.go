package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WordBreak(t *testing.T) {
	testCases := []struct {
		String   string
		Words    []string
		Expected []string
	}{
		{
			"leetcode",
			[]string{"leet", "code"},
			[]string{"leet code"},
		},
		{
			"applepenapple",
			[]string{"apple", "pen"},
			[]string{"apple pen apple"},
		},
		{
			"catsandog",
			[]string{"cats", "oge", "and", "cat"},
			[]string{},
		},
		{
			"pineapplepenapple",
			[]string{"apple", "pen", "applepen", "pine", "pineapple"},
			[]string{"pine apple pen apple", "pine applepen apple", "pineapple pen apple"},
		},
	}
	for _, tc := range testCases {
		actual := wordBreak(tc.String, tc.Words)
		assert.Equal(t, tc.Expected, actual, "Did not work for string"+tc.String)
	}
}
