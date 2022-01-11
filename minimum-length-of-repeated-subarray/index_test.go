package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MinimumAbsoluteDifference(t *testing.T) {
	testCases := []struct {
		Nums1    []int
		Nums2    []int
		Expected int
	}{
		{
			[]int{1, 2, 3, 2, 1},
			[]int{3, 2, 1, 4, 7},
			3,
		},
		{
			[]int{0, 0, 0, 0, 0},
			[]int{0, 0, 0, 0, 0},
			5,
		},
		{
			[]int{0, 0, 0, 0, 0, 2},
			[]int{0, 0, 0, 0, 0, 1},
			5,
		},
		{
			[]int{0, 0, 0, 0, 0, 2, 0},
			[]int{0, 0, 0, 0, 0, 1, 0},
			5,
		},
		{
			[]int{2, 3, 4, 5},
			[]int{0, 0, 0, 0, 0},
			0,
		},
		{
			[]int{1},
			[]int{1},
			1,
		},
	}
	for _, tc := range testCases {
		actual := findLength(tc.Nums1, tc.Nums2)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v & %v", tc.Nums1, tc.Nums2))
	}
}
