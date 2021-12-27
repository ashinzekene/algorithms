package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findPeakElement(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected int
	}{
		{
			[]int{5, 7, 7, 8, 8, 10},
			5,
		},
		{
			[]int{1, 2, 3, 5, 6, 7, 8, 8, 9, 10},
			9,
		},
		{
			[]int{1, 2},
			1,
		},
		{
			[]int{5, 4, 3, 4, 5},
			4, // or 0
		},
		{
			[]int{4, 3, 2},
			0,
		},
	}
	for _, tc := range testCases {
		actual := findPeakElement(tc.Nums)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for %v", tc.Nums))
	}
}
