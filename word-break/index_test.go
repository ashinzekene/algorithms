package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WordBreak(t *testing.T) {
	testCases := []struct {
		String   string
		Words    []string
		Expected bool
	}{
		{
			"leetcode",
			[]string{"leet", "code"},
			true,
		},
		{
			"applepenapple",
			[]string{"apple", "pen"},
			true,
		},
		{
			"catsandog",
			[]string{"cats", "dog", "sand", "and", "cat"},
			false,
		},
		{
			"catsandog",
			[]string{"cats", "oge", "and", "cat"},
			false,
		},
		{
			"catsandog",
			[]string{"cats", "sand", "cat", "og"},
			true,
		},
		{
			"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
			[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
			false,
		},
	}
	for _, tc := range testCases {
		actual := wordBreak(tc.String, tc.Words)
		assert.Equal(t, tc.Expected, actual, "Did not work for string"+tc.String)
	}
}
