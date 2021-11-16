package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetMaximumGenerated(t *testing.T) {
	testCases := []struct {
		N        int
		Expected int
	}{
		{
			7,
			3,
		},
		{
			3,
			2,
		},
		{
			2,
			1,
		},
		{
			1,
			0,
		},
		{
			0,
			0,
		},
		{
			20,
			7,
		},
	}
	for _, tc := range testCases {
		result := GetMaximumGenerated(tc.N)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for %d", tc.N))
	}
}
