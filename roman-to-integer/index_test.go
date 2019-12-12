package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReverseInteger(t *testing.T) {
	testCases := []struct {
		String   string
		Expected int
	}{
		{
			"VIII",
			8,
		},
		{
			"IV",
			4,
		},
		{
			"XV",
			15,
		},
		{
			"XL",
			40,
		},
		{
			"MDCII",
			1602,
		},
		{
			"MCDII",
			1402,
		},
		{
			"MCMXCIV",
			1994,
		},
	}
	for _, tc := range testCases {
		actual := romanToInt(tc.String)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.String))
	}
}
