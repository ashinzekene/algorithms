package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TheLoveLetterMystery(t *testing.T) {
	testCases := []struct {
		String   string
		Expected int32
	}{
		{
			"abc",
			2,
		},
		{
			"abcba",
			0,
		},
		{
			"abcd",
			4,
		},
		{
			"cba",
			2,
		},
	}
	for _, tc := range testCases {
		actual := TheLoveLetterMystery(tc.String)
		assert.Equal(t, tc.Expected, actual)
	}
}
