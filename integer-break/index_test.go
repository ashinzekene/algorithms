package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_integerBreak(t *testing.T) {
	testCases := []struct {
		N        int
		Expected int
	}{
		{
			24,
			6561,
		},
		{
			12,
			81,
		},
		{
			2,
			1,
		},
		{
			16,
			324,
		},
		{
			8,
			18,
		},
		{
			3,
			2,
		},
	}
	for _, tc := range testCases {
		actual := integerBreak(tc.N)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for %d", tc.N))
	}
}
