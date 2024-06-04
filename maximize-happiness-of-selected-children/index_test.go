package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LeftRotation(t *testing.T) {
	testCases := []struct {
		Nums     []int
		K					int
		Expected int64
	}{
		{
			[]int{3, 2, 6, 4},
			2,
			9,
		},
		{
			[]int{1, 2, 3},
			2,
			4,
		},
	}
	for _, tc := range testCases {
		actual := maximumHappinessSum(tc.Nums, tc.K)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Nums))
	}
}
