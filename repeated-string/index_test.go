package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RepeatedString(t *testing.T) {
	testCases := []struct {
		String   string
		Repeats  int64
		Expected int64
	}{
		{
			"aba",
			10,
			7,
		},
		{
			"a",
			1000000000000,
			1000000000000,
		},
		{
			"boo",
			1000000000000,
			0,
		},
	}
	for _, tc := range testCases {
		actual := RepeatedString(tc.String, tc.Repeats)
		assert.Equal(t, tc.Expected, actual)
		actual = RepeatedString1(tc.String, tc.Repeats)
		assert.Equal(t, tc.Expected, actual)
	}
}
