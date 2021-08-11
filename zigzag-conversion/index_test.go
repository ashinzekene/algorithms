package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ZigZagConversion(t *testing.T) {
	testCases := []struct {
		S        string
		NumRows  int
		Expected string
	}{
		{
			"PAYPALISHIRING",
			3,
			"PAHNAPLSIIGYIR",
		},
		{
			"PAYPALISHIRING",
			4,
			"PINALSIGYAHRPI",
		},
		{
			"AB",
			1,
			"AB",
		},
		{
			"A",
			1,
			"A",
		},
		{
			"AB",
			2,
			"AB",
		},
	}
	for _, tc := range testCases {
		actual := convert(tc.S, tc.NumRows)
		assert.Equal(t, tc.Expected, actual, "Did not work for "+tc.S)
	}
}
