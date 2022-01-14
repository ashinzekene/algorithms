package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_monotoneIncreasingDigits(t *testing.T) {
	testCases := []struct {
		N        int
		Expected int
	}{
		{
			10,
			9,
		},
		{
			19,
			19,
		},
		{
			12221,
			11999,
		},
		{
			12341234,
			12339999,
		},
		{
			4321,
			3999,
		},
	}
	for _, tc := range testCases {
		actual := monotoneIncreasingDigits(tc.N)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %d", tc.N))
	}

}
