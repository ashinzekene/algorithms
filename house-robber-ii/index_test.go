package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_rob(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected int
	}{
		{
			[]int{1, 2, 3, 1},
			4,
		},
		{
			[]int{2, 7, 9, 3, 1},
			11,
		},
		{
			[]int{2, 1, 1, 2},
			3,
		},
		{
			[]int{2, 1},
			2,
		},
		{
			[]int{1, 2},
			2,
		},
		{
			[]int{1, 2, 3},
			3,
		},
		{
			[]int{2, 3, 2},
			3,
		},
	}
	for _, tc := range testCases {
		actual := rob(tc.Nums)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for groupAnagrams1 using: %v", tc.Nums))
	}
}
