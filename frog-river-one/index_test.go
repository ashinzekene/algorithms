package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FrogRiverOne(t *testing.T) {
	testCases := []struct {
		Nums     []int
		X        int
		Expected int
	}{
		{
			[]int{1, 3, 1, 4, 2, 3, 5, 4},
			5,
			6,
		},
		{
			[]int{1, 2, 3},
			3,
			2,
		},
		{
			[]int{1, 2, 2, 2, 3},
			4,
			-1,
		},
	}
	for _, tc := range testCases {
		actual := FrogRiverOne(tc.X, tc.Nums)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for nums: %v and X: %d", tc.Nums, tc.X))
	}
}
