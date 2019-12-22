package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LargestRectangle(t *testing.T) {
	testCases := []struct {
		String   string
		Expected []string
	}{
		{
			"23",
			[]string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			"12",
			[]string{},
		},
		{
			"02",
			[]string{},
		},
		{
			"01",
			[]string{},
		},
	}
	for _, tc := range testCases {
		actual := letterCombinations(tc.String)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.String))
	}
}
