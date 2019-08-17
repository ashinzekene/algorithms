package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValid(t *testing.T) {
	testCases := []struct {
		String   string
		Expected string
	}{
		{
			"aabbcd",
			"NO",
		},
		{
			"abcdefghhgfedecba",
			"YES",
		},
		{
			"aabbccddeefghi",
			"NO",
		},
	}
	for _, tc := range testCases {
		actual := isValid(tc.String)
		assert.Equal(t, tc.Expected, actual)
	}
}
