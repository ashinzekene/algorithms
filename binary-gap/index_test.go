package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_binaryGap(t *testing.T) {
	testCases := []struct {
		Num      int
		Expected int
	}{
		{
			5,
			1,
		},
		{
			1041,
			5,
		},
		{
			1,
			0,
		},
		{
			20,
			1,
		},
		{
			1,
			0,
		},
		{
			6291457,
			20,
		},
		{
			66561,
			9,
		},
		{
			1024,
			0,
		},
	}
	for _, tc := range testCases {
		result := binaryGap(tc.Num)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for %v", tc.Num))
	}
}
