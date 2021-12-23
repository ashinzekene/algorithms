package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_titleToNumber(t *testing.T) {
	testCases := []struct {
		N        int
		Expected int
	}{
		{
			1,
			0,
		},
		{
			12,
			2,
		},
		{
			26,
			6,
		},
		{
			126,
			31,
		},
		{
			635,
			158,
		},
		{
			2132,
			531,
		},
		{
			3126,
			781,
		},
	}
	for _, tc := range testCases {
		actual := trailingZeroes(tc.N)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for %d ", tc.N))
	}
}
