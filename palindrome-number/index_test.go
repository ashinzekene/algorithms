package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PalindromeNumber(t *testing.T) {
	testCases := []struct {
		Number   int
		Expected bool
	}{
		{
			365723,
			false,
		},
		{
			36563,
			true,
		},
		{
			6556,
			true,
		},
		{
			-434,
			false,
		},
		{
			1234567654321,
			true,
		},
		{
			8,
			true,
		},
	}
	for _, tc := range testCases {
		actual := isPalindrome(tc.Number)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Number))
	}
}
