package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_intToRoman(t *testing.T) {
	testCases := []struct {
		Number   int
		Expected string
	}{
		{
			256,
			"CCLVI",
		},
		{
			26,
			"XXVI",
		},
		{
			3026,
			"MMMXXVI",
		},
		{
			3999,
			"MMMCMXCIX",
		},
		{
			3909,
			"MMMCMIX",
		},
		{
			109,
			"CIX",
		},
		{
			9,
			"IX",
		},
	}
	for _, tc := range testCases {
		actual := intToRoman(tc.Number)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for Cost %d", tc.Number))
	}
}
