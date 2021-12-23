package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getMoneyAmount(t *testing.T) {
	testCases := []struct {
		N        int
		Expected int
	}{
		{
			1,
			0,
		},
		{
			2,
			1,
		},
		{
			4,
			4,
		},
		{
			5,
			6,
		},
		{
			8,
			12,
		},
		{
			10,
			16,
		},
	}
	for _, tc := range testCases {
		actual := getMoneyAmount(tc.N)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for %d ", tc.N))
	}
}
