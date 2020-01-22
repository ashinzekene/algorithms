package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TheLoveLetterMystery(t *testing.T) {
	testCases := []struct {
		Number   int
		Expected int
	}{
		{
			12,
			3,
		},
		{
			100,
			10,
		},
		{
			1,
			1,
		},
		{
			600,
			24,
		},
	}
	for _, tc := range testCases {
		actual := mySqrt(tc.Number)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for %v", tc.Number))
	}
}
