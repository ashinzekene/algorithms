package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_titleToNumber(t *testing.T) {
	testCases := []struct {
		Str      string
		Expected int
	}{
		{
			"AZ",
			52,
		},
		{
			"XB",
			626,
		},
		{
			"A",
			1,
		},
		{
			"Z",
			26,
		},
	}
	for _, tc := range testCases {
		actual := titleToNumber(tc.Str)
		assert.Equal(t, tc.Expected, actual, "Failed for List "+tc.Str)
	}
}
