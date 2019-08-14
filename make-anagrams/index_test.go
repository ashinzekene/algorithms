package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LeftRotation(t *testing.T) {
	testCases := []struct {
		StringA string
		StringB string
		Expected  int32
	}{
		{
			"cde",
			"abc",
			4,
		},
		{
			"fcrxzwscanmligyxyvym",
			"jxwtrhvujlmrpdoqbisbwhmgpmeoke",
			30,
		},
		{
			"showman",
			"woman",
			2,
		},
	}
	for _, tc := range testCases {
		actual := MakeAnagram(tc.StringA, tc.StringB)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not wortk for %s and %s", tc.StringA, tc.StringB))
	}
}
