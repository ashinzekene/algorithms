package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MinDistance(t *testing.T) {
	testCases := []struct {
		Word1    string
		Word2    string
		Expected int
	}{
		{
			"horse",
			"ros",
			3,
		},
		{
			"intention",
			"execution",
			5,
		},
		{
			"goat",
			"goat",
			0,
		},
		{
			"goa",
			"goat",
			1,
		},
		{
			"g",
			"",
			1,
		},
		{
			"",
			"",
			0,
		},
		{
			"a",
			"ab",
			1,
		},
	}
	for _, tc := range testCases {
		result := MinDistance(tc.Word1, tc.Word2)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for pages %s and %s", tc.Word1, tc.Word2))
	}
}
