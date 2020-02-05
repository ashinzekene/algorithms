package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DecodeWays(t *testing.T) {
	testCases := []struct {
		String   string
		Expected int
	}{
		{
			"1262",
			3,
		},
		{
			"1262",
			3,
		},
		{
			"26",
			2,
		},
		{
			"127",
			2,
		},
		{
			"1270",
			0,
		},
		{
			"83778549129",
			2,
		},
		{
			"8254779486",
			2,
		},
		{
			"122231131122",
			120,
		},
		{
			"122212313113",
			126,
		},
		{
			"321121311231",
			65,
		},
	}
	for _, tc := range testCases {
		result := numDecodings(tc.String)
		assert.Equal(t, tc.Expected, result, "Failed for %v", tc.String)
	}
}
