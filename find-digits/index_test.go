package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindDigits(t *testing.T) {
	testCases := []struct {
		Digit    int32
		Expected int32
	}{
		{
			12,
			2,
		},
		{
			1012,
			3,
		},
		{
			1043,
			1,
		},
		{
			123456789,
			3,
		},
		{
			114108089,
			3,
		},
		{
			110110015,
			6,
		},
		{
			121,
			2,
		},
		{
			33,
			2,
		},
		{
			44,
			2,
		},
		{
			55,
			2,
		},
		{
			66,
			2,
		},
		{
			77,
			2,
		},
		{
			88,
			2,
		},
		{
			106108048,
			5,
		},
	}
	for _, tc := range testCases {
		actual := FindDigits(tc.Digit)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for Digit %d", tc.Digit))
	}
}
