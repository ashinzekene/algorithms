package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		Clouds   []int32
		Expected int32
	}{
		{
			[]int32{0, 0, 1, 0, 0, 1, 0},
			4,
		},
		{
			[]int32{0, 0, 0, 1, 0, 0},
			3,
		},
	}
	for _, tc := range testCases {
		actual := JumpingOnClouds(tc.Clouds)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for %v", tc.Clouds))
	}
}
