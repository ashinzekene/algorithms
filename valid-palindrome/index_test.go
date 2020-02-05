package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ValidPalindrome(t *testing.T) {
	testCases := []struct {
		String   string
		Expected bool
	}{
		{
			"365723",
			false,
		},
		{
			"A man, a plan, a canal: Panama",
			true,
		},
		{
			"s'os",
			true,
		},
		{
			"pas))(*&))ap",
			true,
		},
		{
			")P",
			true,
		},
		{
			"0P",
			false,
		},
		{
			"12321",
			true,
		},
	}
	for _, tc := range testCases {
		actual := isPalindrome(tc.String)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.String))
	}
}
