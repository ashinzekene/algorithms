package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MinCostClimbingStairs(t *testing.T) {
	testCases := []struct {
		Costs    []int
		Expected int
	}{
		{
			[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			6,
		},
		{
			[]int{0, 0, 0, 0},
			0,
		},
		{
			[]int{0, 1, 0, 0},
			0,
		},
		{
			[]int{10, 15, 20},
			15,
		},
	}
	for _, tc := range testCases {
		actual := minCostClimbingStairs(tc.Costs)
		assert.Equal(t, tc.Expected, actual)
	}
}
