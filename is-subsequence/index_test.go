package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MinCostClimbingStairs(t *testing.T) {
	testCases := []struct {
		S        string
		T        string
		Expected bool
	}{
		{
			"abc",
			"ahbgdc",
			true,
		},
		{
			"axc",
			"ahbgdc",
			false,
		},
	}
	for _, tc := range testCases {
		actual := isSubsequence(tc.S, tc.T)
		assert.Equal(t, tc.Expected, actual)
	}
}
