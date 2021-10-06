package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_climbingArea(t *testing.T) {
	testCases := []struct {
		A        int
		B        int
		K        int
		Expected int
	}{
		{
			6,
			11,
			2,
			3,
		},
		{
			2,
			2,
			2,
			1,
		},
		{
			0,
			12,
			3,
			4,
		},
	}
	for _, tc := range testCases {
		result := CountDiv(tc.A, tc.B, tc.K)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for A: %v, B: %v, K: %v", tc.A, tc.B, tc.K))
	}
}
