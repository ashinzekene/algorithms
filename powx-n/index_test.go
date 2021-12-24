package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MyPow(t *testing.T) {
	testCases := []struct {
		X        float64
		N        int
		Expected float64
	}{
		{
			2,
			10,
			1024,
		},
		{
			2,
			1,
			2,
		},
		{
			2,
			-1,
			0.5,
		},
		{
			2,
			-2,
			0.25,
		},
		{
			2,
			0,
			1,
		},
	}
	for _, tc := range testCases {
		actual := myPow(tc.X, tc.N)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("plusOne did not work for X: %v and power N: %v", tc.X, tc.N))
	}
}
