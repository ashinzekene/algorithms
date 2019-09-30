package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MinimumAbsoluteDifference(t *testing.T) {
	testCases := []struct {
		Array    []int32
		Expected int32
	}{
		{
			[]int32{3, -7, 0},
			3,
		},
		{
			[]int32{-59, -36, -13, 1, -53, -92, -2, -96, -54, 75},
			1,
		},
		{
			[]int32{1, -3, 71, 68, 17},
			3,
		},
	}
	for _, tc := range testCases {
		actual := minimumAbsoluteDifference(tc.Array)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Array))
	}
}
