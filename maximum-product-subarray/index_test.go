package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LeftRotation(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected int
	}{
		{
			[]int{2, 3, -2, 4},
			6,
		},
		{
			[]int{-2, 3, -4},
			24,
		},
		{
			[]int{-2, 0, -1},
			0,
		},
	}
	for _, tc := range testCases {
		actual := maxProduct(tc.Nums)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Nums))
	}
}
