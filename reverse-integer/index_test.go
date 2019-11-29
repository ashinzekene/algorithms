package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReverseInteger(t *testing.T) {
	testCases := []struct {
		Integer  int
		Expected int
	}{
		{
			1000,
			1,
		},
		{
			-42345,
			-54324,
		},
		{
			12345,
			54321,
		},
		{
			122387378278399,
			0,
		},
	}
	for _, tc := range testCases {
		actual := reverse(tc.Integer)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Integer))
	}
}
