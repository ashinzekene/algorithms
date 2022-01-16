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
			12,
		},
		{
			[]int{2, 1, 1, 2},
			4,
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
			[]int{1, 1, 1, 2},
			3,
		},
	}
	for _, tc := range testCases {
		actual := rob(tc.Nums)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("rob failed for groupAnagrams1 using: %v", tc.Nums))
		actual = rob1(tc.Nums)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("rob1 failed for groupAnagrams1 using: %v", tc.Nums))
	}
}
