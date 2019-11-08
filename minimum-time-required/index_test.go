package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MinimumTimeRequired(t *testing.T) {
	testCases := []struct {
		Machines []int64
		Goal     int64
		Expected int64
	}{
		{
			[]int64{1, 3, 4},
			10,
			7,
		},
		{
			[]int64{4, 5, 6},
			12,
			20,
		},
		{
			[]int64{2, 3, 2},
			10,
			8,
		},
		{
			[]int64{2, 3},
			5,
			6,
		},
		{
			[]int64{477, 448, 240, 858, 927, 703, 172, 68, 969, 943, 657, 499, 753, 777, 438, 199, 356, 435, 63, 292, 446, 164, 323, 511, 145, 607, 39, 832, 764, 692, 990, 240, 491, 581, 98, 769, 635, 621, 189, 603, 915, 197, 453, 667, 973, 890, 865, 328, 676, 928},
			306,
			1904,
		},
	}
	for _, tc := range testCases {
		actual := minTime(tc.Machines, tc.Goal)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Machines))
	}

}
