package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AlternatingCharacters(t *testing.T) {
	testCases := []struct {
		String   string
		Expected int32
	}{
		{
			"AAAA",
			3,
		},
		{
			"BBBBB",
			4,
		},
		{
			"ABABABAB",
			0,
		},
		{
			"BABABA",
			0,
		},
		{
			"AAABBB",
			4,
		},
	}
	for _, tc := range testCases {
		result := AlternatingCharacters(tc.String)
		assert.Equal(t, tc.Expected, result, "Failed for "+tc.String)
	}
}
