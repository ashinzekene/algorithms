package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_multiplyStrings(t *testing.T) {
	testCases := []struct {
		Num1     string
		Num2     string
		Expected string
	}{
		{
			"12",
			"24",
			"288",
		},
		{
			"123",
			"456",
			"56088",
		},
		{
			"498828660196",
			"840477629533",
			"419254329864656431168468",
		},
		{
			"984594",
			"0",
			"0",
		},
	}
	for _, tc := range testCases {
		actual := multiply(tc.Num1, tc.Num2)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v * %v", tc.Num1, tc.Num2))
	}

}
