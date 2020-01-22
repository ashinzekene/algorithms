package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxArea(t *testing.T) {
	testCases := []struct {
		Count    int
		Expected string
	}{
		{
			1,
			"1",
		},
		{
			2,
			"11",
		},
		{
			3,
			"21",
		},
		{
			4,
			"1211",
		},
		{
			5,
			"111221",
		},
		{
			6,
			"312211",
		},
		{
			7,
			"13112221",
		},
		{
			8,
			"1113213211",
		},
		{
			9,
			"31131211131221",
		},
		{
			10,
			"13211311123113112211",
		},
	}
	for _, tc := range testCases {
		result := countAndSay(tc.Count)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for %v", tc.Count))
	}
}
