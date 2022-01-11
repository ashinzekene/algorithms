package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MinimumAbsoluteDifference(t *testing.T) {
	testCases := []struct {
		Nums     []int
		X        int
		Expected int
	}{
		{
			[]int{1, 1, 4, 2, 3},
			5,
			2,
		},
		{
			[]int{3, 2, 20, 1, 1, 3},
			10,
			5,
		},
		{
			[]int{8, 5, 6, 4, 5, 2},
			4,
			-1,
		},
		{
			[]int{8828, 9581, 49, 9818, 9974, 9869, 9991, 10000, 10000, 10000, 9999, 9993, 9904, 8819, 1231, 6309},
			134365,
			16,
		},
	}
	for _, tc := range testCases {
		actual := minOperations(tc.Nums, tc.X)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Nums))
	}
}
