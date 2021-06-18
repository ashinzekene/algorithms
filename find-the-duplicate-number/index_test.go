package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findTheDuplicateNumber(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected int
	}{
		{
			[]int{1, 1},
			1,
		},
		{
			[]int{1, 2, 1},
			1,
		},
		{
			[]int{3, 4, 2, 3},
			3,
		},
		{
			[]int{3, 1, 3, 4, 2},
			3,
		},
		{
			[]int{2, 5, 7, 8, 6, 1, 3, 4, 6},
			6,
		},
		{
			[]int{1, 1, 2},
			1,
		},
	}
	for _, tc := range testCases {
		actual := findDuplicate1(tc.Nums)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed findDuplicate for %v", tc.Nums))
		actual = findDuplicate2(tc.Nums)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed findDuplicate2 for %v", tc.Nums))
		actual = findDuplicate(tc.Nums)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed findDuplicate1 for %v", tc.Nums))
	}
}
