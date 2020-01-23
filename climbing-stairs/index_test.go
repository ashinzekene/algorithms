package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxArea(t *testing.T) {
	testCases := []struct {
		Num      int
		Expected int
	}{
		{
			1,
			1,
		},
		{
			2,
			2,
		},
		{
			3,
			3,
		},
		{
			4,
			5,
		},
		{
			5,
			8,
		},
		{
			12,
			233,
		},
	}
	for _, tc := range testCases {
		result := climbStairs(tc.Num)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for %v", tc.Num))
	}
}
