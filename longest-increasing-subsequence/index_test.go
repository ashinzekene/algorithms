package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLIS(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected int
	}{
		{
			[]int{10, 9, 2, 5, 3, 7, 101, 18},
			4,
		},
		{
			[]int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15},
			6,
		},
		{
			[]int{3, 10, 2, 1, 20},
			3,
		},
		{
			[]int{3, 2},
			1,
		},
		{
			[]int{50, 3, 10, 7, 40, 80},
			4,
		},
		{
			[]int{50, 53, 55, 5, 6, 7, 40, 80},
			5,
		},
	}
	for _, tc := range testCases {
		actual := lengthOfLIS(tc.Nums)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("lengthOfLIS did not work for %v", tc.Nums))
	}
}
