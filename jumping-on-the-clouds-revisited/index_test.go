package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_JumpingOnTheCloudsRevisited(t *testing.T) {
	testCases := []struct {
		Clouds   []int32
		JumpSize int32
		Expected int32
	}{
		{
			[]int32{0, 0, 1, 0, 0, 1, 1, 0},
			2,
			92,
		},
		{
			[]int32{0, 0, 1, 0},
			2,
			96,
		},
		{
			[]int32{1, 1, 1, 0, 1, 1, 0, 0, 0, 0},
			3,
			94,
		},
	}
	for _, tc := range testCases {
		actual := jumpingOnClouds(tc.Clouds, tc.JumpSize)
		assert.Equal(t, tc.Expected, actual)
	}
}
