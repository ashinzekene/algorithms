package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMedianSortedArrays(t *testing.T) {
	testCases := []struct {
		Nums1     []int
		Nums2     []int
		Expected  float64
	}{
		{
			[]int{1, 2},
			[]int{3, 4},
			2.5,
		},
		{
			[]int{1, 2},
			[]int{3},
			2,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{8},
			3.5,
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{},
			3,
		},
	}
	for _, tc := range testCases {
		actual := findMedianSortedArrays(tc.Nums1, tc.Nums2)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v and %v", tc.Nums1, tc.Nums2))
	}
}
