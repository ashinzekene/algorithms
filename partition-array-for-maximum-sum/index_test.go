package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MaxSumAfterPartitioning(t *testing.T) {
	testCases := []struct {
		Arr      []int
		K        int
		Expected int
	}{
		{
			[]int{1, 15, 7, 9, 2, 5, 10},
			3,
			84,
		},
		{
			[]int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3},
			4,
			83,
		},
		{
			[]int{8, 5, 7},
			3,
			24,
		},
		{
			[]int{3, 5, 7},
			3,
			21,
		},
		{
			[]int{1},
			1,
			1,
		},
	}
	for _, tc := range testCases {
		actual := MaxSumAfterPartitioning(tc.Arr, tc.K)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Partition did not work for %v", tc.Arr))
	}
}
