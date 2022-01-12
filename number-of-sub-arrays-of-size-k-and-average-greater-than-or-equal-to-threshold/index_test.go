package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numOfSubarrays(t *testing.T) {
	testCases := []struct {
		Nums      []int
		K         int
		Threshold int
		Expected  int
	}{
		{
			[]int{2, 2, 2, 2, 5, 5, 5, 8},
			3,
			4,
			3,
		},
		{
			[]int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3},
			3,
			5,
			6,
		},
		{
			[]int{1, 3, 4, 3},
			3,
			3,
			1,
		},
		{
			[]int{1, 3, 5, 3},
			3,
			2,
			2,
		},
		{
			[]int{1, 3, 5, 3},
			4,
			3,
			1,
		},
		{
			[]int{1, 3, 5, 3},
			4,
			4,
			0,
		},
	}
	for _, tc := range testCases {
		actual := numOfSubarrays(tc.Nums, tc.K, tc.Threshold)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v, K: %d, Threshold: %d", tc.Nums, tc.K, tc.Threshold))
	}
}
