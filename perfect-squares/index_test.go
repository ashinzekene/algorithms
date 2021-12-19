package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumSquares(t *testing.T) {
	testCases := []struct {
		N        int
		Expected int
	}{
		{
			12,
			3,
		},
		{
			13,
			2,
		},
		{
			120,
			3,
		},
		{
			1,
			1,
		},
		{
			7,
			4,
		},
	}
	for _, tc := range testCases {
		actual := NumSquares(tc.N)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.N))
	}
}
