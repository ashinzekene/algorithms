package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getPermutation(t *testing.T) {
	testCases := []struct {
		N        int
		K        int
		Expected string
	}{
		{
			3,
			5,
			"312",
		},
		{
			3,
			3,
			"213",
		},
		{
			4,
			9,
			"2314",
		},
		{
			6,
			7,
			"124356",
		},
	}
	for _, tc := range testCases {
		actual := getPermutation(tc.N, tc.K)
		assert.Equal(t, tc.Expected, actual)
	}
}
