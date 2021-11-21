package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LongestPalindrome(t *testing.T) {
	testCases := []struct {
		String   string
		Expected string
	}{
		{
			"babad",
			"bab",
		},
		{
			"a",
			"a",
		},
		{
			"aba",
			"aba",
		},
		{
			"abacabcaroline",
			"bacab",
		},
		{
			"abaccabcaroline",
			"baccab",
		},
	}
	for _, tc := range testCases {
		actual := LongestPalindrome(tc.String)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("LongestPalindrome did not work for %s", tc.String))
	}
}
