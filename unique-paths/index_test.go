package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TrapTainWater(t *testing.T) {
	testCases := []struct {
		M        int
		N        int
		Expected int
	}{
		{
			3,
			2,
			3,
		},
		{
			3,
			3,
			6,
		},
		{
			3,
			6,
			21,
		},
		{
			4,
			4,
			20,
		},
		{
			7,
			4,
			84,
		},
		{
			5,
			4,
			35,
		},
	}
	for _, tc := range testCases {
		actual := uniquePaths(tc.M, tc.N)
		assert.Equal(t, tc.Expected, actual)
	}
}
