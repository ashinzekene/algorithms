package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_3SumClosest(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Target   int
		Expected int
	}{
		{
			[]int{-1, 0, 1, 2, -1, -4},
			4,
			3,
		},
		{
			[]int{6, -2, 4, 7, 5},
			11,
			11,
		},
		{
			[]int{1, -1, -1, 0, 2, -3, 2, 4},
			4,
			4,
		},
	}
	for _, tc := range testCases {
		actual := threeSumClosest(tc.Nums, tc.Target)
		assert.Equal(t, tc.Expected, actual)
	}
}
